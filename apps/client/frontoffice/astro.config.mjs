import {defineConfig} from 'astro/config';
import tailwind from '@astrojs/tailwind';

// https://astro.build/config
import node from "@astrojs/node";

// https://astro.build/config
import sitemap from "@astrojs/sitemap";

export default defineConfig({
    outDir: '../../dist/apps/frontoffice',
    integrations: [
        tailwind({

            config: {
                applyBaseStyles: true
            }
        }),
        sitemap({
            customPages: [
            ],
            i18n: {
                defaultLocale: 'en',
                locales:{
                    en: 'en-US'
                }
            }
        })
    ],
    output: 'server',
    vite: {
        ssr: {
            // Example: Force a broken package to skip SSR processing, if needed
            noExternal: ['date-fns']
        }
    },
    adapter: node({
        mode: 'standalone'
    })
});
