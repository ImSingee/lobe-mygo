import { afterEach, expect, test } from 'bun:test';

import { createReloadServer, type ReloadServer } from './reload-server';

let server: ReloadServer | undefined;

afterEach(() => {
  server?.stop();
  server = undefined;
});

const nextEvents = async (reader: ReadableStreamDefaultReader<Uint8Array>, count: number) => {
  const decoder = new TextDecoder();
  let text = '';
  const events: unknown[] = [];
  while (events.length < count) {
    const { done, value } = await reader.read();
    if (done) break;
    text += decoder.decode(value, { stream: true });
    let end = text.indexOf('\n\n');
    while (end !== -1) {
      const chunk = text.slice(0, end);
      text = text.slice(end + 2);
      const data = chunk
        .split('\n')
        .filter((line) => line.startsWith('data: '))
        .map((line) => line.slice(6))
        .join('\n');
      if (data) events.push(JSON.parse(data));
      end = text.indexOf('\n\n');
    }
  }
  return events;
};

const connect = async (port: number) => {
  const res = await fetch(`http://localhost:${port}/events`);
  return { reader: res.body!.getReader(), res };
};

test('serves events with CORS and an event-stream type', async () => {
  server = createReloadServer(0);
  const { res, reader } = await connect(server.port);
  expect(res.headers.get('content-type')).toContain('text/event-stream');
  expect(res.headers.get('access-control-allow-origin')).toBe('*');
  reader.cancel();
});

test('sends the current state to a new client, then broadcasts changes', async () => {
  server = createReloadServer(0);
  server.broadcast({ id: 'a1', type: 'built' });
  const { reader } = await connect(server.port);
  server.broadcast({ type: 'building' });
  server.broadcast({ message: 'button.go:3: undefined: x', type: 'error' });
  expect(await nextEvents(reader, 3)).toEqual([
    { id: 'a1', type: 'built' },
    { type: 'building' },
    { message: 'button.go:3: undefined: x', type: 'error' },
  ]);
  reader.cancel();
});

test('a client joining after a failed build sees the error', async () => {
  server = createReloadServer(0);
  server.broadcast({ id: 'a1', type: 'built' });
  server.broadcast({ message: 'boom', type: 'error' });
  const { reader } = await connect(server.port);
  expect(await nextEvents(reader, 1)).toEqual([{ message: 'boom', type: 'error' }]);
  reader.cancel();
});
