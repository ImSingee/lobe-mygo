# lobe-mygo web catalog: design

## Intent

lobe-mygo is a component library for MyGo native UI (`github.com/egoist/mygo/ui`) that reproduces the
look of `@lobehub/ui`. Its components draw on the GPU, so nobody can see them without running Go
code. The catalog fixes that: a docs website, laid out like the lobe-ui docs (docs-kit), whose
component pages run the real Go components live in the browser through MyGo's WebAssembly + WebGL 2
backend.

The catalog has two jobs:

1. **Acceptance surface.** From now on every component is checked in the browser here (agent-browser
   against the dev server), side by side with the matching page on ui.lobehub.com, in light and dark.
2. **Docs.** The same pages are what users of lobe-mygo read: demos, the Go source of each demo, and
   a props table.

## Goals

- `bun run dev` serves a docs-kit site; editing a Go demo or a component and reloading the page shows
  the change (no manual `go build`).
- `bun run build` produces a static site with the WebAssembly included.
- A component page shows one or more live demos, each with its Go source below it, and an API table
  generated from the Go code.
- The demos follow the site's light/dark toggle.
- First pages: the Button page (current `lobe.Button`) and a Design Tokens page that draws the
  generated token swatches with MyGo (checks the token codegen visually).

## Non-goals

- New components or new Button features (Danger, Ghost, Shape, Loading…). The catalog ships with what
  exists today.
- Deployment (Vercel / Cloudflare), custom domain, analytics, giscus.
- Accessibility of the canvas content: MyGo's browser backend exposes nothing to assistive
  technology yet (upstream limitation).
- Geist font. The web build gets text from the browser, so Geist would be possible there, but desktop
  would differ; deferred.
- Porting any lobe-ui docs prose.

## Anti-goals

- Do not reimplement components in React/HTML to make the docs "look right". Every pixel inside a demo
  frame comes from the Go code that desktop apps run.
- Do not fork or patch docs-kit. Use it as a consumer, through `docs.config.ts` and MDX only.

## Constraints (verified)

- **MyGo in the browser** (`docs/ui/browsers.md` upstream): `GOOS=js GOARCH=wasm go build`, loaded with
  Go's `wasm_exec.js`; a window fills the element with id `mygo`, else the page. One instance shows one
  window, so each demo runs in its own `<iframe>`. Size about 9–10 MB raw, 2.6 MB gzip for a build
  with every widget; `-tags mygo_noinspector -trimpath -ldflags="-s -w"` trims it.
- **Theme in a frame**: a frame's `prefers-color-scheme` follows the browser, not the page.
  MyGo reads `data-mygo-theme="dark" | "light"` on the frame document's root element and follows it
  live.
- **docs-kit** (`@lobehub/docs-kit` 6.3.1, CLI `lobedocs`): the page's appearance is
  `document.documentElement.dataset.theme` (`'dark' | 'light'`). Pages are `index.mdx` files under the
  `atomDirs` roots; URL from the folder (`catalog/button/index.mdx` → `/components/button`). `public/`
  is served at `/`. Only `Demo` and `Api` are MDX globals; other components are imported. **docs-kit
  loads its own `vite.config.ts` and takes no consumer Vite plugins**, so the WebAssembly cannot be
  built by a Vite plugin the way MyGo's own website does it (`website/docs-demo.ts` upstream); it has
  to be built by a script into `public/`.
- docs-kit's `<Demo of={X}>` only accepts `?demo` imports of a TSX file with one default export, and
  `<Api>` only reads TypeScript. Neither fits Go demos, so the catalog brings its own two MDX
  components (below). `lobedocs build` fails on a `<Demo>` / `<Api>` that does not resolve, so the
  catalog pages must not use them.
- Precedent to copy: MyGo's website (`website/docs-demo/`, `website/docs-demo.ts`, upstream
  `egoist/mygo`). Its Go app registers demos by name in a map, picks one from `os.Args`, and its loader
  page (`index.html`) passes `?demo=` as `go.argv`, syncs the parent's theme into `data-mygo-theme`
  with a MutationObserver, and shows a spinner and a load error message. Port those parts; they are
  proven.

## Decided

### Repository layout (all inside `~/git/work/lobe-mygo`)

```
catalog/
  main.go                  package main: registry lookup, preview frame, os.Args → demo
  registry/registry.go     Add(name, view), Get(name), Names()
  button/
    index.mdx              the docs page (/components/button)
    basic.go               one demo per file, registered in init()
    variants.go
    sizes.go
  tokens/
    index.mdx              Design Tokens page (/components/tokens)
    swatches.go
docs/
  index.mdx                home (frontmatter only; default docs-kit home)
  components/GoDemo.tsx    MDX component: live frame + Go source
  components/GoApi.tsx     MDX component: props table from generated JSON
web/mygo/index.html        loader page template (ported from MyGo's docs-demo/index.html)
scripts/
  build-wasm.ts            go build → public/mygo/{catalog.wasm, wasm_exec.js, index.html}; --watch
  gen-api.go (or cmd/)     go/doc over package lobe → catalog/<name>/api.json
docs.config.ts
public/mygo/               build output, gitignored
```

- Demo packages live **next to their `index.mdx`**, one idea per file, as lobe-ui does with
  `demos/<Idea>.tsx`. Demo names are `<page>/<file>` (`button/variants`).
- **Demo file convention (my call):** a demo file holds only `package`, its imports and one exported
  view function (`func Variants(c *ui.Context)`), so the file shown under the frame is what a user
  copies. Registration lives in one `register.go` per package (`init()` calling `registry.Add` for
  each demo); `catalog/main.go` imports the packages with blank imports. Demo state lives in
  package-level variables of the demo file.
- `atomDirs: [{ dir: 'catalog' }]`. Frontmatter as docs-kit requires (`title`, `description`,
  `category`); categories reuse lobe-ui's names (`General` for Button).

### Demo frame (`catalog/main.go`)

- `go run ./catalog button/variants` opens that demo in a desktop window; no argument opens an index of
  all demos (as MyGo's docs-demo does). The web loader passes the same argv.
- The preview wraps the demo: `lobe.Apply(c)`, then a column that fills the window, centered, padding
  20, gap 12, background **`BgContainer`** of the current lobe tokens (lobe-ui's docs show demos on
  the container surface, not the layout gray).
- To read the current tokens from outside the package, export the existing helper:
  `tokensFor` → `lobe.TokensOf(c *ui.Context) *Tokens`.

### Loader page (`web/mygo/index.html` → `public/mygo/index.html`)

- `/mygo/?demo=button/variants`. Instantiates `/mygo/catalog.wasm` with `wasm_exec.js`, `go.argv =
  ["catalog", demo]`.
- Theme sync: read `parent.document.documentElement.dataset.theme`, write `data-mygo-theme`, observe
  the parent's `data-theme` attribute with a MutationObserver (docs-kit's attribute, not MyGo
  website's `class="dark"`).
- Spinner while loading; on failure show the error text in the frame.
- Body background transparent until the canvas paints.

### `<GoDemo>` (`docs/components/GoDemo.tsx`)

```mdx
import { GoDemo } from '../../docs/components/GoDemo';
import variants from './variants.go?raw';

<GoDemo name="button/variants" title="Variants" height={160} source={variants} />
```

- Renders a card in the docs-kit demo style: title, the iframe (`loading="lazy"`, fixed `height`,
  full width, no border), and below it the Go source, collapsed by default, with a copy button.
- Lazy frames matter: each frame instantiates its own ~10 MB module; a page should keep to about
  four demos.

### `<GoApi>` (`docs/components/GoApi.tsx`)

- `scripts/gen-api` parses package `lobe` with `go/doc`, and for each exported `XxxProps` struct writes
  its fields: name, Go type, doc comment, and for enum types (`ButtonVariant`, `ButtonSize`) the
  constant names in order, the first being the zero-value default.
- Output: `catalog/<page>/api.json`, gitignored, regenerated by `dev` and `build`.
- `<GoApi data={api} name="ButtonProps" />` renders the table: Property, Type, Default, Description.
- This needs doc comments on the exported fields of `ButtonProps` and on the enum constants. They are
  public library API docs, the one kind of comment the project allows.

### Scripts (`package.json`)

- `dev`: run `gen-api`, then `build-wasm --watch` and `lobedocs dev` concurrently. The watcher rebuilds
  the wasm when any `.go` file outside `node_modules` changes (debounced ~200 ms) and regenerates
  the API JSON; reloading the page picks up the new module (the loader fetches with `cache: 'no-store'`
  in dev).
- `build`: `gen-api`, `build-wasm` (production flags, content-hashed file name written into the loader
  page), then `lobedocs build`.
- `docs.config.ts`: `title: 'Lobe MyGo'`, `atomDirs`, `prefersColor: 'auto'`, `siteUrl` placeholder.

### Design Tokens page (my call)

- One demo, `tokens/swatches`, in a frame tall enough to scroll inside (MyGo `ui.Scroll`), following
  the site theme like every demo.
- Semantic groups only, in this order: Text, Background, Fill, Border, Primary, Error, Warning,
  Success, Info. Palette scales (`Red1`…) are left out.
- `gen-tokens.ts` also emits `var colorTokenNames = []string{...}` and
  `func (t *Tokens) Color(name string) ui.Color` (a generated switch), so the page lists tokens
  without reflection. Each swatch shows the color, its name and its hex/rgba value.
- Pass check: values match ui.lobehub.com/design-tokens for the same names in light and dark.

### Acceptance

The builder proves it in a browser (agent-browser, against `bun run dev`):

1. `/components/button` shows three live frames; each canvas has painted (not just the spinner).
2. Toggling the site theme turns the frames dark within one second without reloading them.
3. Hovering a Filled button inside a frame changes its fill (screenshot before/after).
4. The API table lists `Variant` and `Size` with defaults `Default` and `Middle`.
5. Editing `catalog/button/variants.go` and reloading shows the edit.
6. `bun run build` succeeds and the built site loads the hashed wasm.

7. `/components/tokens` scrolls through every semantic group; spot-check five values against
   ui.lobehub.com/design-tokens in both themes.

Screenshots of the Button page in light and dark next to ui.lobehub.com/components/button go in the
report. The live site runs `@lobehub/ui` 6.15.2, the same version the tokens come from, so it is the
reference.

**Mismatches (my call):** the catalog passes when it runs live and correctly; it does not have to look
identical. Differences found in the comparison go in the report as a list. The builder fixes one only
when it is a wrong token or metric in `button.go` / `gen-tokens.ts` that takes under 15 minutes;
anything larger is a follow-up.

## Left to the builder

- Concurrency tool for `dev` (`concurrently`, a small Bun script, …).
- Code highlighting for the Go source: reuse what docs-kit / `@lobehub/ui` exports (`Highlighter`)
  if it renders Go; otherwise plain `<pre>`.
- Card styling details of `<GoDemo>`; matching docs-kit's own demo card is the target.
- Swatch layout of the tokens page.
- Whether `gen-api` is `scripts/gen-api.go` run with `go run` or a `cmd/` package.

## Artifacts

None. The prior research and the prototype live in the conversation and in this repo
(`button.go`, `theme.go`, `tokens_gen.go`, `examples/gallery`).
