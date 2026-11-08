export type ReloadEvent =
  | { id: string; type: 'built' }
  | { type: 'building' }
  | { message: string; type: 'error' };

export interface ReloadServer {
  broadcast: (event: ReloadEvent) => void;
  port: number;
  stop: () => void;
}

const encoder = new TextEncoder();
const frame = (event: ReloadEvent) => encoder.encode(`data: ${JSON.stringify(event)}\n\n`);

export function createReloadServer(port: number): ReloadServer {
  const clients = new Set<ReadableStreamDefaultController<Uint8Array>>();
  let last: ReloadEvent | undefined;

  const send = (client: ReadableStreamDefaultController<Uint8Array>, chunk: Uint8Array) => {
    try {
      client.enqueue(chunk);
    } catch {
      clients.delete(client);
    }
  };

  const heartbeat = setInterval(() => {
    for (const client of clients) send(client, encoder.encode(': ping\n\n'));
  }, 15_000);

  const server = Bun.serve({
    fetch(req) {
      if (new URL(req.url).pathname !== '/events') return new Response('Not found', { status: 404 });
      let controller: ReadableStreamDefaultController<Uint8Array>;
      const stream = new ReadableStream<Uint8Array>({
        cancel() {
          clients.delete(controller);
        },
        start(c) {
          controller = c;
          clients.add(c);
          send(c, encoder.encode(': connected\n\n'));
          if (last) send(c, frame(last));
        },
      });
      return new Response(stream, {
        headers: {
          'Access-Control-Allow-Origin': '*',
          'Cache-Control': 'no-store',
          'Content-Type': 'text/event-stream',
        },
      });
    },
    idleTimeout: 0,
    port,
  });

  return {
    broadcast(event) {
      last = event;
      const chunk = frame(event);
      for (const client of clients) send(client, chunk);
    },
    port: server.port ?? port,
    stop() {
      clearInterval(heartbeat);
      for (const client of clients) {
        try {
          client.close();
        } catch {}
      }
      clients.clear();
      server.stop(true);
    },
  };
}
