import '@fortawesome/fontawesome-free/css/all.min.css'
import type {User} from "$lib/types/user";
import {redirect} from "@sveltejs/kit";

export const prerender: boolean = true
export const ssr: boolean = false

/** @type {import('./$types').PageLoad} */
export async function load(): Promise<void> {
    let user:User|null
    const u:string|null = localStorage.getItem('auth_user')
    user = !!u ? JSON.parse(u) : null

    if(user !== null && !!user.username) {
        redirect(302, "/opportunities")
    }
}
