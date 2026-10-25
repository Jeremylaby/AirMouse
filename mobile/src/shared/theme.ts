export const colors = {
  background: '#FFFFFF',
} as const;

export const spacing = {
  xs: 4,
  sm: 8,
  md: 16,
  lg: 24,
  xl: 32,
} as const;

export const radius = {
  sm: 4,
  md: 8,
  lg: 16,
  full: 9999,
} as const;

export const fontSize = {
  sm: 12,
  md: 16,
  lg: 20,
  xl: 28,
} as const;

export const theme = { colors, spacing, radius, fontSize } as const;

export type Theme = typeof theme;
