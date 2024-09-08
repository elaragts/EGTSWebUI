import {defineStore} from 'pinia';
import type { Item } from '@/types';

export const useDatatableStore = defineStore('datatable', {
    state: () => ({
        initialized: false,
        headList: [] as Item[],
        bodyList: [] as Item[],
        faceList: [] as Item[],
        kigurumiList: [] as Item[],
        puchiList: [] as Item[],
        titleList: [] as Item[],
        songMap: {}
    }),
    actions: {
        async initDatatable(): Promise<void> {
            if (this.initialized) return; // Prevent re-initialization

            const response = await fetch('/api/datatable')
            if (response.status === 200) {
                const res = await response.json();

                this.headList = res.head;
                this.bodyList = res.body;
                this.faceList = res.face;
                this.kigurumiList = res.kigurumi;
                this.puchiList = res.puchi;
                this.titleList = res.title;
                this.songMap = res.song;
            } else {
                console.error('Unexpected response status:', response.status);
            }

            this.initialized = true; // Mark as initialized
        }
    }

});