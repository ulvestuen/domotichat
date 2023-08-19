<script>
    export let role = 'user';
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
        max-width: 480px;
        text-align: left;
    }

    .user {
        background-color: #ffd653;
        align-self: flex-end;
    }

    .assistant {
        background-color: #5cc07c;
        align-self: flex-start;
    }
</style>
