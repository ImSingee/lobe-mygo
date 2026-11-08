import { BentoCard, BentoGrid } from '@lobehub/ui/awesome';

import { renderLink } from './renderLink';
import { Shot } from './Shot';

const TILES = [
  { colSpan: 2, href: '/components/agent-avatar', hint: 'seeded shapes · animated states', name: 'agents', rowSpan: 2, title: 'AgentAvatar', width: 520 },
  { colSpan: 2, href: '/components/highlighter', hint: 'syntax highlighting · copy', name: 'code', rowSpan: 2, title: 'Highlighter', width: 520 },
  { colSpan: 1, href: '/components/date-picker', hint: 'keyboard navigation', name: 'calendar', rowSpan: 2, title: 'Calendar', width: 260 },
  { colSpan: 1, href: '/components/qr-code', hint: 'icon · statuses', name: 'qr', rowSpan: 2, title: 'QRCode', width: 260 },
  { colSpan: 1, href: '/components/switch', hint: 'switch · slider · rate', name: 'controls', rowSpan: 2, title: 'Controls', width: 260 },
  { colSpan: 1, href: '/components/button', hint: 'variants · sizes', name: 'buttons', rowSpan: 1, title: 'Button', width: 260 },
  { colSpan: 1, href: '/components/group-avatar', hint: 'grids · corners', name: 'group', rowSpan: 1, title: 'GroupAvatar', width: 260 },
  { colSpan: 2, href: '/components/alert', hint: 'status colors', name: 'alerts', rowSpan: 1, title: 'Alert', width: 520 },
  { colSpan: 2, href: '/components/color-swatches', hint: 'Lobe UI palette', name: 'swatches', rowSpan: 1, title: 'ColorSwatches · Tag', width: 520 },
] as const;

export function Gallery() {
  return (
    <BentoGrid>
      {TILES.map((tile) => (
        <BentoCard
          colSpan={tile.colSpan}
          hint={tile.hint}
          href={tile.href}
          key={tile.name}
          renderLink={renderLink}
          rowSpan={tile.rowSpan}
          title={tile.title}
        >
          <Shot alt={`${tile.title} rendered by Lobe MyGo`} name={tile.name} width={tile.width} />
        </BentoCard>
      ))}
    </BentoGrid>
  );
}
