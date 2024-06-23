export const toUpperCase = (s: string) => s.toUpperCase()

export const parseAmount = (s: string) => Number(s.replaceAll(",", "."))
