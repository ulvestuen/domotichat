<script>
    export let role = 'user'; // 'user' or 'gpt-4'
    export let text = '';
</script>

<div class="message {role}">
    {text}
</div>

<style>
    .message {
        padding: 10px;
        margin: 5px;
        border-radius: 5px;
    }

    .message.user {
        background-color: #f3f3f3;
        align-self: flex-end;
    }

    .message.gpt-4 {
        background-color: #e0e7ff;
        align-self: flex-start;
    }
</style>
