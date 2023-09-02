<script>
    import Message from './Message.svelte';
    import '@authorizerdev/authorizer-svelte/styles/default.css';

    let input = "";
    let messages = [];

    let sendIconColor = "#ff9c00"

    async function sendMessage() {
        if (!input) {
            return
        }
        const userMessage = {
            role: 'user',
            content: input
        };
        messages = [...messages, userMessage];
        input = "";

        messages = await postChatCompletions(messages)
    }

    async function postChatCompletions(messages) {
        const endpoint = import.meta.env.VITE_DOMOTICHAT_API_BASE_URL + "/chat/completions"
        console.log(endpoint)
        const response = await fetch(endpoint, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(messages)
        });
        return await response.json()
    }

</script>

<main>
    <div class="chat-box">
        {#each messages as message}
            {#if (["user", "assistant"].findIndex(v => v === message.role) > -1)
            && message.content }
                <Message role={message.role} text={message.content}/>
            {/if}
        {/each}
    </div>
    <div class="input-area">
        <input bind:value={input}
               on:keydown={e => e.key === "Enter" && sendMessage()}
               class="text-input"/>
        <button on:click={sendMessage}
                on:mouseenter={() => {sendIconColor = "#ee9300"}}
                on:mouseleave={() => {sendIconColor = "#ff9c00"}}
                class="send-button">
            <svg xmlns="http://www.w3.org/2000/svg"
                 viewBox="0 0 24 24"
                 width="24"
                 height="24">
                <path stroke={sendIconColor}
                      fill={sendIconColor}
                      d="M1.513 1.96a1.374 1.374 0 0 1 1.499-.21l19.335 9.215a1.147 1.147 0 0 1 0 2.07L3.012 22.25a1.374 1.374 0 0 1-1.947-1.46L2.49 12 1.065 3.21a1.375 1.375 0 0 1 .448-1.25Zm2.375 10.79-1.304 8.042L21.031 12 2.584 3.208l1.304 8.042h7.362a.75.75 0 0 1 0 1.5Z"></path>
            </svg>
        </button>
    </div>

</main>

<style>
    .chat-box {
        display: flex;
        flex-direction: column;
        width: 80vw;
    }

    .input-area {
        display: flex;
        flex-direction: row;
        width: 80vw;
        align-items: center;
        justify-content: flex-end;
        margin-top: 20px;
        padding: 4px 4px 4px 4px;
        background-color: #ffd653;
        border-radius: 8px;
        box-shadow: 0 0 10px rgb(0 0 0 / 0.2);
    }

    .text-input {
        flex-grow: 1;
        margin-right: 10px;
        padding: 0.6em;
        font-family: Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;
        font-size: 1.2em;
        color: #ff9c00;
        border: 1px solid #ffd653;
        border-radius: 8px;
        background-color: #ffd653;
        box-shadow: 0 0 10px rgb(0 0 0 / 0.2);
        outline: none;
    }

    .send-button {
        display: flex;
        place-items: center center;
        color: #ff9c00;
        background-color: #ffd653;
        border-style: none;
        margin: 0 8px 0 8px;
    }

    .send-button:hover {
        cursor: pointer;
        color: bisque;
    }

</style>
