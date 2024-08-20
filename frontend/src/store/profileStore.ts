import {defineStore} from 'pinia';
import {useAuthStore} from "@/store/authStore";
import {useToast} from 'vue-toastification';
import type { ProfileOptions, CostumeOptions, SongOptions } from '@/types';

export const useProfileStore = defineStore('profile', {
    state: () => ({
        initialized: false,
        profileOptions: {} as ProfileOptions,
        costumeOptions: {} as CostumeOptions,
        songOptions: {} as SongOptions
    }),
    actions: {
        async initProfile(baid: number): Promise<void> {
            if (this.initialized) return; // Prevent re-initialization

            const profile = await fetch(`/api/user/${baid}/profile_options`)
            if (profile.status === 200) {
                this.profileOptions = await profile.json();
            } else {
                console.error('Unexpected user profile data response status:', profile.status);
            }

            const costume = await fetch(`/api/user/${baid}/costume_options`)
            if (costume.status === 200) {
                this.costumeOptions = await costume.json();
            } else {
                console.error('Unexpected user costume data response status:', costume.status);
            }

            const song = await fetch(`/api/user/${baid}/song_options`)
            if (song.status === 200) {
                this.songOptions = await song.json();
            } else {
                console.error('Unexpected user song data response status:', song.status);
            }

            this.initialized = true; // Mark as initialized
        },
        async uploadProfile(): Promise<boolean> {
            const data = {
                profileOptions: this.profileOptions,
                costumeOptions: this.costumeOptions,
                songOptions: this.songOptions
            }

            const authStore = useAuthStore(); // for user baid
            const response = await fetch(`/api/user/${authStore.baid}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });

            const toast = useToast();
            if (response.ok) {
                toast.success('Profile Updated Successfully');
                return true;
            } else {
                toast.error(`Failed to update profile: ${await response.text()}`);
                return false;
            }
        }
    }

});