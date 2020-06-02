const polarToCartesian = (centerX, centerY, radius, angleInDegrees)  => {
  const angleInRadians = (angleInDegrees - 90) * Math.PI / 180

  return {
    x: centerX + radius * Math.cos(angleInRadians),
    y: centerY + radius * Math.sin(angleInRadians)
  };
}

export const arcPath = (x, y, radius, startAngle, endAngle) => {
    const start = polarToCartesian(x, y, radius, startAngle)
    const end = polarToCartesian(x, y, radius, endAngle)
    const largeArcFlag = startAngle - endAngle <= 180 ? "0" : "1"

    return [
        "M", start.x, start.y,
        "A", radius, radius, 0, largeArcFlag, 0, end.x, end.y
    ].join(" ")
}
