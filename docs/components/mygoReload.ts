import { useSyncExternalStore } from 'react';

export interface ReloadState {
  building: boolean;
  error: string | null;
  version: number;
}

type ReloadEvent = { id: string; type: 'built' } | { type: 'building' } | { message: string; type: 'error' };

const RELOAD_URL = 'http://localhost:5199/events';
const initial: ReloadState = { building: false, error: null, version: 0 };

let state = initial;
let source: EventSource | null = null;
let loadedId: string | undefined;
let heard = false;
const listeners = new Set<() => void>();

const set = (next: Partial<ReloadState>) => {
  state = { ...state, ...next };
  for (const listener of listeners) listener();
};

const onEvent = (event: ReloadEvent) => {
  switch (event.type) {
    case 'building': {
      set({ building: true });
      break;
    }
    case 'error': {
      set({ building: false, error: event.message });
      break;
    }
    case 'built': {
      const first = !heard;
      if (first) loadedId = event.id;
      if (event.id === loadedId) {
        set({ building: false, error: null });
      } else {
        loadedId = event.id;
        set({ building: false, error: null, version: state.version + 1 });
      }
      break;
    }
  }
  heard = true;
};

const connect = () => {
  if (source || !import.meta.env.DEV || typeof EventSource === 'undefined') return;
  source = new EventSource(RELOAD_URL);
  source.addEventListener('message', (message) => onEvent(JSON.parse(message.data) as ReloadEvent));
};

const subscribe = (listener: () => void) => {
  connect();
  listeners.add(listener);
  return () => listeners.delete(listener);
};

export const useMygoReload = () =>
  useSyncExternalStore(
    subscribe,
    () => state,
    () => initial,
  );
