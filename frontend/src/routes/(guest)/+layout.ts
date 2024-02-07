import type {PageLoadEvent} from "../../../.svelte-kit/types/src/routes/$types";
import type {User} from "$lib/types/user";
import {redirect} from "@sveltejs/kit";

export const prerender: boolean = true
export const ssr: boolean = false

/** @type {import('./$types').PageLoad} */
export async function load(pageEvent: PageLoadEvent): Promise<void> {
    let user: User | null
    const u: string | null = localStorage.getItem('auth_user')
    user = !!u ? JSON.parse(u) : null

    if (!!user) {
        redirect(302, "/")
    }
}