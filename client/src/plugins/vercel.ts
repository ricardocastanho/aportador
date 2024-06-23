import { inject } from '@vercel/analytics'

export const loadVercel = () => {
  const mode = import.meta.env.PROD ? 'production' : 'development'

  inject({
    mode: mode,
  })
}
