interface ShotProps {
  alt: string;
  name: string;
  width: number;
}

export function Shot({ alt, name, width }: ShotProps) {
  return (
    <span className="mygo-shot" style={{ maxWidth: width }}>
      <img alt={alt} className="mygo-shot-light" decoding="async" loading="lazy" src={`/home/${name}-light.png`} />
      <img alt="" aria-hidden className="mygo-shot-dark" decoding="async" loading="lazy" src={`/home/${name}-dark.png`} />
    </span>
  );
}
