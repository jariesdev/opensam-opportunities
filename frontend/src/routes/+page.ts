// since there's no dynamic data here, we can prerender
// it so that it gets served as a static asset in production
import type {User} from "$lib/types/user";
import {redirect} from "@sveltejs/kit";
import type {PageLoadEvent} from "../../.svelte-kit/types/src/routes/$types";

export const prerender = true;

/** @type {import('./$types').PageLoad} */
export async function load(pageEvent: PageLoadEvent): Promise<void> {
    let user: User | null
    const u: string | null = localStorage.getItem('auth_user')
    user = !!u ? JSON.parse(u) : null

    if (pageEvent.route.id === '/') {
        if (user !== null && !!user.username) {
            redirect(302, "/opportunities")
        } else {
            redirect(302, "/login")
        }
    }
}