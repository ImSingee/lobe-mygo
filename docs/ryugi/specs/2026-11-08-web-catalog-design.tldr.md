# TL;DR: lobe-mygo web catalog

**What gets built:** a docs website laid out like ui.lobehub.com (docs-kit), whose component pages run
the real Go components live in the browser (MyGo → WebAssembly → WebGL 2). It is where every
component gets checked from now on.

## What you'll see

1. `bun run dev` → open `/components/button`.
2. Three live demo frames (Basic, Variants, Sizes), each with its Go source below, collapsed.
3. An API table for `ButtonProps` generated from the Go code.
4. Toggle the site theme → the demos turn dark without reloading.
5. Edit a `.go` demo, reload → the change shows.
6. `/components/tokens` draws every lobe token swatch with MyGo.

## Decisions

| Topic | Decision |
| --- | --- |
| Docs framework | docs-kit as a consumer, no fork |
| How Go runs on the page | one `<iframe>` per demo, `/mygo/?demo=button/variants` |
| WASM build | script into `public/mygo/` + watcher (docs-kit takes no Vite plugins) |
| Demo files | `catalog/<page>/<idea>.go` next to `index.mdx`, one idea per file |
| Demo source display | `import x from './x.go?raw'` into a `<GoDemo>` component |
| API table | `go/doc` → `api.json` → `<GoApi>` |
| Theme sync | parent `data-theme` → frame `data-mygo-theme`, live |

## Check these (my calls)

- Demos sit on `BgContainer` (white), like lobe-ui's docs, not the window gray.
- `tokensFor` becomes public as `lobe.TokensOf(c)`.
- Doc comments get added to `ButtonProps` fields and enum constants, to feed the API table
  (public library API, the allowed kind of comment).
- A Design Tokens page is in scope (semantic groups only, no palette scales); new Button features are not.
- A demo file = one exported view function, no registry code, so the shown source is copyable;
  registration goes in a separate `register.go`.
- Pass = runs live and correctly. Visual differences from ui.lobehub.com get listed; only quick
  token/metric fixes (under 15 min) get done in this round.
- No deployment yet.

## Time

About 3–4 hours of build plus 30 minutes of browser acceptance.
