/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'

import { inject } from "@vercel/analytics";

const app = createApp(App)

registerPlugins(app)

const mode = import.meta.env.PROD ? "production" : "development";

inject({
  mode: mode,
});

app.mount('#app')
