<script>
    import {AuthorizerProvider} from "@authorizerdev/authorizer-svelte";
    import '@authorizerdev/authorizer-svelte/styles/default.css';
    import AuthWrapper from "./lib/AuthWrapper.svelte";
</script>

<AuthorizerProvider config={{
        authorizerURL: "https://authorizer-production-7757.up.railway.app",
        redirectURL: typeof window != "undefined" ? window.location.origin : "",
        client_id: "c8f1f70d-0e6c-4b02-a282-8d5197d30648"
    }}>
    <AuthWrapper/>
</AuthorizerProvider>

<style>
</style>
