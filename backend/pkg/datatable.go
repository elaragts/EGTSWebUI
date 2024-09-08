package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type dataPairs struct {
	Id           uint   `json:"id"`
	EnglishName  string `json:"englishName"`
	JapaneseName string `json:"japaneseName"`
}

type songPairs struct {
	EnglishName  string `json:"englishName"`
	JapaneseName string `json:"japaneseName"`
}

type Data struct {
	Head     []dataPairs        `json:"head"`
	Body     []dataPairs        `json:"body"`
	Face     []dataPairs        `json:"face"`
	Kigurumi []dataPairs        `json:"kigurumi"`
	Puchi    []dataPairs        `json:"puchi"`
	Title    []dataPairs        `json:"title"`
	Song     map[uint]songPairs `json:"song"`
}

type wordlistItem struct {
	Key               string `json:"key"`
	JapaneseText      string `json:"japaneseText"`
	JapaneseFontType  uint   `json:"japaneseFontType"`
	EnglishUsText     string `json:"englishUsText"`
	EnglishUsFontType uint   `json:"englishUsFontType"`
	ChineseTText      string `json:"chineseTText"`
	ChineseTFontType  uint   `json:"chineseTFontType"`
	KoreanText        string `json:"koreanText"`
	KoreanFontType    uint   `json:"koreanFontType"`
	ChineseSText      string `json:"chineseSText"`
	ChineseSFontType  uint   `json:"chineseSFontType"`
}

type wordlistType struct {
	Items []wordlistItem `json:"items"`
}

type costumeItem struct {
	CosType  string `json:"cosType"`
	UniqueID uint   `json:"uniqueId"`
	DataID   uint   `json:"dataId"`
}

type costumeType struct {
	Items []costumeItem `json:"items"`
}

type shougouItem struct {
	UniqueID uint `json:"uniqueId"`
	Rarity   uint `json:"rarity"`
}

type shougouType struct {
	Items []shougouItem `json:"items"`
}

type songItem struct {
	SongId   string `json:"id"`
	UniqueID uint   `json:"uniqueId"`
}

type songType struct {
	Items []songItem `json:"items"`
}

var Datatable Data

func InitDatatable(datatablePath string) {
	costumesFile, err := os.Open(datatablePath + "/don_cos_reward.json")
	if err != nil {
		fmt.Println("Error opening costumes file:", err)
		return
	}
	wordlistFile, err := os.Open(datatablePath + "/wordlist.json")
	if err != nil {
		fmt.Println("Error opening wordlist file:", err)
		return
	}
	shougouFile, err := os.Open(datatablePath + "/shougou.json")
	if err != nil {
		fmt.Println("Error opening shougou file:", err)
		return
	}
	songsFile, err := os.Open(datatablePath + "/musicinfo.json")
	if err != nil {
		fmt.Println("Error opening music file:", err)
		return
	}
	defer costumesFile.Close()
	defer wordlistFile.Close()
	defer shougouFile.Close()
	defer songsFile.Close()

	var wordlist wordlistType

	decoder := json.NewDecoder(wordlistFile)
	err = decoder.Decode(&wordlist)
	if err != nil {
		fmt.Println("Error decoding wordlist file:", err)
		return
	}

	// need to loop through wordlist file and put all data points
	// that are either a costume or title into a hashmap
	// so that it is more efficient to perform look-ups when reading
	// the other two files

	// extra type just to make hashmap values easier to read
	type stringPair struct {
		EnglishUsText string `json:"englishUsText"`
		JapaneseText  string `json:"japaneseText"`
	}

	wordlistMap := make(map[string]stringPair)
	for _, item := range wordlist.Items {
		if !strings.HasPrefix(item.Key, "costume") &&
			!strings.HasPrefix(item.Key, "syougou") &&
			(!strings.HasPrefix(item.Key, "song") ||
				strings.HasPrefix(item.Key, "song_sub") ||
				strings.HasPrefix(item.Key, "song_detail")) {

			// filtering out anything that isn't a valid costume, syougou, or song
			// also filters out song_detail and song_sub as they aren't needed
			continue
		}

		var pair stringPair
		pair.JapaneseText = item.JapaneseText
		pair.EnglishUsText = item.EnglishUsText
		wordlistMap[item.Key] = pair
	}

	var costumes costumeType

	decoder = json.NewDecoder(costumesFile)
	err = decoder.Decode(&costumes)
	if err != nil {
		fmt.Println("Error decoding costumes file:", err)
		return
	}

	// need to loop through costumes file
	// take each id, look it up in the wordlistMap
	// depending on type of costume: head, body, face, kigurumi, puchi
	// we send the data to a different array in the Datatable variable,
	// each data point should be sent as a dataPairs type

	for _, item := range costumes.Items {
		switch item.CosType {
		case "head":
			var newHeadItem dataPairs
			newHeadItem.Id = item.UniqueID

			// key in the wordlist map is in the format costume_costumetype_uniqueid
			var key string
			key = fmt.Sprintf("%s%d", "costume_head_", item.UniqueID)

			newHeadItem.EnglishName = wordlistMap[key].EnglishUsText
			newHeadItem.JapaneseName = wordlistMap[key].JapaneseText

			Datatable.Head = append(Datatable.Head, newHeadItem)
		case "body":
			var newBodyItem dataPairs
			newBodyItem.Id = item.UniqueID

			// key in the wordlist map is in the format costume_costumetype_uniqueid
			var key string
			key = fmt.Sprintf("%s%d", "costume_body_", item.UniqueID)

			newBodyItem.EnglishName = wordlistMap[key].EnglishUsText
			newBodyItem.JapaneseName = wordlistMap[key].JapaneseText

			Datatable.Body = append(Datatable.Body, newBodyItem)
		case "face":
			var newFaceItem dataPairs
			newFaceItem.Id = item.UniqueID

			// key in the wordlist map is in the format costume_costumetype_uniqueid
			var key string
			key = fmt.Sprintf("%s%d", "costume_face_", item.UniqueID)

			newFaceItem.EnglishName = wordlistMap[key].EnglishUsText
			newFaceItem.JapaneseName = wordlistMap[key].JapaneseText

			Datatable.Face = append(Datatable.Face, newFaceItem)
		case "kigurumi":
			var newKigurumiItem dataPairs
			newKigurumiItem.Id = item.UniqueID

			// key in the wordlist map is in the format costume_costumetype_uniqueid
			var key string
			key = fmt.Sprintf("%s%d", "costume_kigurumi_", item.UniqueID)

			newKigurumiItem.EnglishName = wordlistMap[key].EnglishUsText
			newKigurumiItem.JapaneseName = wordlistMap[key].JapaneseText

			Datatable.Kigurumi = append(Datatable.Kigurumi, newKigurumiItem)
		case "puchi":
			var newPuchiItem dataPairs
			newPuchiItem.Id = item.UniqueID

			// key in the wordlist map is in the format costume_costumetype_uniqueid
			var key string
			key = fmt.Sprintf("%s%d", "costume_puchi_", item.UniqueID)

			newPuchiItem.EnglishName = wordlistMap[key].EnglishUsText
			newPuchiItem.JapaneseName = wordlistMap[key].JapaneseText

			Datatable.Puchi = append(Datatable.Puchi, newPuchiItem)
		}
	}

	var shougou shougouType

	decoder = json.NewDecoder(shougouFile)
	err = decoder.Decode(&shougou)
	if err != nil {
		fmt.Println("Error decoding shougou file:", err)
		return
	}

	// need to loop through shougou file
	// take each id, look it up in the wordlistMap
	// we send the data to the title array in the Datatable variable,
	// each data point should be sent as a dataPairs type

	for _, item := range shougou.Items {
		var newTitleItem dataPairs
		newTitleItem.Id = item.UniqueID

		// key in the wordlist map is in the format syougou_uniqueid
		var key string
		key = fmt.Sprintf("%s%d", "syougou_", item.UniqueID)

		newTitleItem.EnglishName = wordlistMap[key].EnglishUsText
		newTitleItem.JapaneseName = wordlistMap[key].JapaneseText

		Datatable.Title = append(Datatable.Title, newTitleItem)
	}

	var songs songType

	decoder = json.NewDecoder(songsFile)
	err = decoder.Decode(&songs)
	if err != nil {
		fmt.Println("Error decoding songs file:", err)
		return
	}

	// need to loop through music file
	// take each id, which is a string (songId) instead of a number (uniqueId) like the last two files,
	// and look it up in the wordlistMap, grab the name and send the data to the
	// song array in the Datatable variable,
	// each data point should be sent as a dataPairs type

	Datatable.Song = make(map[uint]songPairs)

	for _, item := range songs.Items {
		if item.UniqueID == 0 { // not including tmap4
			continue
		}

		var newSongItem songPairs

		// key in the wordlist map is in the format song_songID
		var key string
		key = fmt.Sprintf("%s%s", "song_", item.SongId)

		newSongItem.EnglishName = wordlistMap[key].EnglishUsText
		newSongItem.JapaneseName = wordlistMap[key].JapaneseText

		Datatable.Song[item.UniqueID] = newSongItem
	}
}
