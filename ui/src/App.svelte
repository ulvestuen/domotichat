<script>
    import Message from './lib/Message.svelte';

    let input = "";
    let messages = [];

    async function sendMessage() {
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
        <textarea bind:value={input}
                  on:keydown={e => e.key === "Enter" && sendMessage()}
                  class="text-input"></textarea>
        <button on:click={sendMessage}>Send</button>
    </div>
</main>

<style>
    .chat-box {
        display: flex;
        flex-direction: column;
        width: 640px;
    }

    .input-area {
        display: flex;
        flex-direction: row;
        width: 640px;
        align-items: center;
        justify-content: flex-end;
        margin-top: 20px;
        padding: 4px 4px 4px 4px;
        background-color: #ffd653;
        border-radius: 8px;
    }

    .text-input {
        flex-grow: 1;
        margin-right: 10px;
        padding: 0.6em;
        font-size: 1.2em;
        border: 1px solid #ffd653;
        border-radius: 8px;
        background-color: #ffd653;
        resize: none;
    }

</style>
