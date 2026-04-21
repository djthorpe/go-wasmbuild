export interface AuthOptions {
    baseUrl?: string | URL;
}

export interface PublicClientConfiguration {
    issuer: string;
    client_id?: string;
}

export type PublicClientConfigurations = Record<string, PublicClientConfiguration>;

export class Auth {
    readonly baseUrl: URL;

    constructor(options: AuthOptions = {}) {
        const baseUrl = options.baseUrl ?? "/";
        const resolved = baseUrl instanceof URL
            ? new URL(baseUrl.toString())
            : new URL(baseUrl, globalThis.location?.origin ?? "http://localhost");
        if (!resolved.pathname.endsWith("/")) {
            resolved.pathname = `${resolved.pathname}/`;
        }
        this.baseUrl = resolved;
    }

    url(path: string, params?: Record<string, string | number | boolean | undefined>): URL {
        const normalizedPath = path.replace(/^\/+/, "");
        const url = new URL(normalizedPath, this.baseUrl);
        if (params) {
            for (const [key, value] of Object.entries(params)) {
                if (value !== undefined) {
                    url.searchParams.set(key, String(value));
                }
            }
        }
        return url;
    }

    authorizeUrl(provider: string, redirectUri: string, state?: string): URL {
        return this.url("auth/authorize", {
            provider,
            redirect_uri: redirectUri,
            state,
        });
    }

    async config(signal?: AbortSignal): Promise<PublicClientConfigurations> {
        const response = await fetch(this.url("config"), {
            headers: {
                Accept: "application/json",
            },
            signal,
        });
        if (!response.ok) {
            throw new Error(`auth config request failed with status ${response.status}`);
        }
        return response.json() as Promise<PublicClientConfigurations>;
    }
}

export default Auth;