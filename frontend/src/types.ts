export type Item = {
    id: number;
    englishName: string;
    japaneseName: string;
}

export type ProfileOptions = {
    myDonName: string;
    title: string;
    language: number;
    titlePlateId: number;
    displayAchievement: boolean;
    achievementDisplayDifficulty: number;
    displayDan: boolean;
    difficultySettingCourse: number;
    difficultySettingStar: number;
    difficultySettingSort: number;
    customTitleOn: boolean;
}

export type CostumeOptions = {
    currentBody: number;
    currentFace: number;
    currentHead: number;
    currentKigurumi: number;
    currentPuchi: number;
    colorBody: number;
    colorFace: number;
    colorLimb: number;
}

export type SongOptions = {
    speedId: number;
    isVanishOn: boolean;
    isInverseOn: boolean;
    randomId: number;
    isSkipOn: boolean;
    isVoiceOn: boolean;
    selectedToneId: number;
    notesPosition: number;
}