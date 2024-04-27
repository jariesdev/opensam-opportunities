<script lang="ts">
    import {GetAllSettings, UpdateSettings} from "$lib/wailsjs/go/main/App";
    import {createEventDispatcher, onMount} from "svelte";
    import {main, database} from "$lib/wailsjs/go/models";
    import {find} from "lodash";

    const dispatch = createEventDispatcher();

    let apiKey: string = ''
    let isProcessing: boolean = false
    let isSuccess: boolean = false

    let resultText: string
    let settings: database.Setting[] = []

    function updateSettings(): void {
        isProcessing = true

        const request: main.SettingRequest[] = [
            {
                key: 'api_key',
                value: apiKey
            }
        ]

        UpdateSettings(request)
            .then(({message, result}) => {
                isSuccess = result
                resultText = message

                dispatch('success', {isSuccess, resultText, apiKey})
            })
            .catch(e => {
                alert(e.message)
            })
            .finally(() => {
                isProcessing = false
            })
    }

    /**
     * Return setting value
     *
     * @param key
     */
    function getSettingValue(key: string): string {
        const found = find(settings, (o:database.Setting) => o.key === key)
        return found ? found.value : "";
    }

    function initializeForm() {
        apiKey = getSettingValue('api_key')
    }

    onMount(() => {
        GetAllSettings()
            .then((response) => {
                settings = response
                initializeForm()
            })
            .catch(e => {
                alert(e.message)
            })
    })
</script>

<section>
    <div class="row">
        <div class="col-sm-5">
            <form method="POST" on:submit|preventDefault={updateSettings}>
                <div class="input-box mb-4" id="input">
                    <label for="apiKey" class="block text-gray-700 text-sm font-bold mb-2">
                        API Key
                    </label>
                    <input autocomplete="off"
                           bind:value={apiKey}
                           class="form-control mb-3"
                           id="apiKey"
                           type="text"
                           placeholder="Enter api key"/>
                </div>

                <button type="submit"
                        class="btn btn-primary btn-block d-block px-5"
                        disabled="{isProcessing || isProcessing}">
                    Save Changes
                </button>
            </form>
        </div>
    </div>
</section>

