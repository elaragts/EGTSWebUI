import {defineStore} from 'pinia';
import {useProfileStore} from "@/store/profileStore";

export const useAuthStore = defineStore('auth', {
    state: () => ({
        isAuthenticated: false,
        username: '',
        baid: 0,
        initialized: false,
    }),
    actions: {
        async initAuth(): Promise<void> {
            if (this.initialized) return; // Prevent re-initialization

            const response = await fetch('/auth/session');
            switch (response.status) {
                case 200:
                    const res = await response.json();
                    this.username = res.username;
                    this.isAuthenticated = true;
                    this.baid = res.baid;
                    const profileStore = useProfileStore();
                    await profileStore.initProfile(this.baid); // load up profile data
                    break;
                case 401:
                    this.isAuthenticated = false;
                    break;
                default:
                    console.error('Unexpected response status:', response.status);
            }

            this.initialized = true; // Mark as initialized
        },
        async login(formData: any): Promise<[boolean, string]> {
            const response = await fetch('/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: formData,
            });

            if (response.ok) {
                const body = await response.json();
                this.username = body.username;
                this.baid = body.baid;
                this.isAuthenticated = true;
                const profileStore = useProfileStore();
                await profileStore.initProfile(this.baid);
                return [true, ''];
            } else {
                return [false, await response.text()];
            }
        },
        async logout(): Promise<void> {
            const response = await fetch('/auth/logout', {method: 'POST'});
            if (response.ok) {
                this.isAuthenticated = false;
                this.username = '';
                this.baid = 0;
                const profileStore = useProfileStore();
                profileStore.initialized = false;
            } else {
                console.error('Failed to logout:', response.status);
            }
        }
    }
});
