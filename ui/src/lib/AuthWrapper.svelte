<script>
    import {Authorizer} from "@authorizerdev/authorizer-svelte";
    import '@authorizerdev/authorizer-svelte/styles/default.css';
    import Chat from "./Chat.svelte";
    import {getContext} from "svelte";

    let state;

    const store = getContext('authorizerContext');

    store.subscribe((data) => {
        state = data;
    });

    const logoutHandler = async () => {
        await state.logout();
    };
</script>

{#if state.token}
    <div class="chat-container">
        <Chat/>
        <button class="logout-button"
                on:click={logoutHandler}>Logout
        </button>
    </div>
{:else}
    <div class="login-container">
        <h1>DomotiChat login</h1>
        <Authorizer/>
    </div>
{/if}

<style>
    .chat-container {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    .logout-button {
        margin-top: 32px;
        padding: 4px 16px 4px 16px;
        background-color: #ff9c00;
        color: #ffd653;
        border: 1px solid #ff9c00;
        border-radius: 4px;

    }

    .login-container {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        color: #5cc07c;
        background-color: #f8f8e5;

        --authorizer-primary-color: #ff9c00;
        --authorizer-primary-disabled-color: #ff9c00;
        --authorizer-gray-color: #d1d5db;
        --authorizer-white-color: #ffffff;
        --authorizer-danger-color: #dc2626;
        --authorizer-success-color: #10b981;
        --authorizer-text-color: #374151;
        --authorizer-fonts-font-stack: -apple-system, system-ui, sans-serif;
        --authorizer-fonts-large-text: 18px;
        --authorizer-fonts-medium-text: 14px;
        --authorizer-fonts-small-text: 12px;
        --authorizer-fonts-tiny-text: 10px;
        --authorizer-radius-card: 4px;
        --authorizer-radius-button: 4px;
        --authorizer-radius-input: 4px;
    }

</style>
