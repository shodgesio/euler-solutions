export const areFactors = (n: number) => (a: number[]) => a.reduce((a, b) => a * b) === n;