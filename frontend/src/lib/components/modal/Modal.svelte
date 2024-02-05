<slot name="activator" {show} {hide}></slot>

<div bind:this={modalRef} class="modal fade modal-xl" tabindex="-1">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title">
                    <slot name="title">
                        {title}
                    </slot>
                </h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <slot name="body">
                    <p>Modal body text goes here.</p>
                </slot>
            </div>
            <div class="modal-footer">
                <slot name="footer"></slot>
            </div>
        </div>
    </div>
</div>

<script lang="ts">
    import * as bootstrap from 'bootstrap/dist/js/bootstrap';
    import {onMount} from "svelte";

    export let title: string = "Modal Title"
    let modalRef: HTMLElement
    let modal: any = null
    let open: boolean = false

    export function show(): void {
        modal.show()
    }

    export function hide(): void {
        modal.hide()
    }

    onMount(() => {
        modal = new bootstrap.Modal(modalRef)
    })
</script>