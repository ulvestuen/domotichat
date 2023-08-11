<script>
  import Message from './lib/Message.svelte';

  let input = "";
  let messages = [];

  function sendMessage() {
    const userMessage = {
      role: 'user',
      text: input
    };
    messages = [...messages, userMessage];

    // Make an API call to GPT-4 LLM here
    // and then push the response to the messages array.
    fetchOpenAIAPI(userMessage).then(response => {
      messages = [...messages, response];
    });

    input = '';
  }

  async function fetchOpenAIAPI(messages) {
    const endpoint = "https://api.openai.com/v1/chat/completions";
    const apiKey = "YOUR_OPENAI_API_KEY"; // Never expose this on the client side!

    const response = await fetch(endpoint, {
      method: "POST",
      headers: {
        "Authorization": `Bearer ${apiKey}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        model: "gpt-4.0",
        messages,
      })
    });

    if (!response.ok) {
      throw new Error("Failed to fetch from OpenAI");
    }

    const data = await response.json();
    return data.choices[0].message;
  }

</script>

<main>
  <div class="chat-box">
    {#each messages as message}
      <Message role={message.role} text={message.text} />
    {/each}
  </div>
  <div class="input-area">
    <textarea bind:value={input}></textarea>
    <button on:click={sendMessage}>Send</button>
  </div>
</main>

<style>
  /* Add your styles here */

</style>
