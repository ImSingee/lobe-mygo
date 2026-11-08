import { FeatureGrid, type FeatureGridItem } from '@lobehub/ui/awesome';
import { Cpu, Palette, Shapes, SunMoon, SwatchBook, Waves } from 'lucide-react';

import { renderLink } from './renderLink';

const ITEMS: FeatureGridItem[] = [
  {
    description: 'Buttons, inputs, menus, overlays, chat views and more that look and move like their React counterparts.',
    icon: Palette,
    title: 'Lobe UI, in Go',
  },
  {
    description: 'Colors, radii and type sizes are generated from @lobehub/ui, so a Go app and a web app share one look.',
    href: '/components/tokens',
    icon: SwatchBook,
    title: 'Same tokens as Lobe UI',
  },
  {
    description: 'theme.Apply follows the system appearance; every component reads its tokens from the theme.',
    icon: SunMoon,
    title: 'Light and dark built in',
  },
  {
    description: 'One small binary per platform, cross-compiled from any machine, with no webview to ship.',
    icon: Cpu,
    title: 'Pure Go, no cgo',
  },
  {
    description: 'Springs and shared-element transitions for indicators, overlays, toasts and side navs.',
    href: '/components/motion',
    icon: Waves,
    title: 'Motion included',
  },
  {
    description: 'All 2,300+ Lucide icons as Go functions, drawn as vectors at any size.',
    href: '/components/icon',
    icon: Shapes,
    title: 'Lucide icons',
  },
];

export function Features() {
  return <FeatureGrid items={ITEMS} renderLink={renderLink} />;
}
