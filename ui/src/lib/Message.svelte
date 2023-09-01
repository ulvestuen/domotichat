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
        color: #ff9c00;
        align-self: flex-end;
        box-shadow: 0 0 10px rgb(0 0 0 / 0.2);
    }

    .assistant {
        background-color: #5cc07c;
        color: #f8f8e5;
        align-self: flex-start;
        box-shadow: 0 0 10px rgb(0 0 0 / 0.2);
    }
</style>
