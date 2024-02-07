import { dev } from '$app/environment';

// we don't need any JS on this page, though we'll load
// it in dev so that we get hot module replacement
export const csr: boolean = dev;

// since there's no dynamic data here, we can prerender
// it so that it gets served as a static asset in production
export const prerender: boolean = true;


/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load({ route,data }) {
    console.log(route, data);
}