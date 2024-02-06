import '@fortawesome/fontawesome-free/css/all.min.css'

export const prerender: boolean = true
export const ssr: boolean = false

/** @type {import('./$types').PageLoad} */
export async function load({ route,data }) {
    console.log('auth layout:', route, data);
}