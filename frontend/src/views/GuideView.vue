<script setup lang="ts">

</script>

<template>
    <div class="font-sans py-8 text-gray-200 bg-cl2 rounded my-5 mx-52 px-10">
      <h2 id="egts-setup---connecting-and-using-egts-which-is-a-public-chn-server" class="text-center underline text-2xl font-bold mb-4">
        EGTS SETUP - Connecting and using EGTS, which is a public JP/39.06 server
      </h2>
      <p class="mb-4">
        <strong class="underline">HOW TO MIGRATE FROM CHN TO 39.06</strong>
        <br/>
        Copy your access code from the EGTS webui <strong>WITHOUT</strong> the BNTTCNID
        <br/>
        Inside card.ini(39.06) replace accessCode1= with your card number (E.G. accessCode1=42429075045032907255)
        <br/>
        Repeat for accessCode2 <strong>IF</strong> needed
      </p>
      <p class="mb-4">
        <strong class="underline">TO USE EGTS, YOU MUST DOWNLOAD THE EGTS PACK FROM THE LINK DOWN BELOW.</strong>
        <br/>
        <a href="https://discord.com/channels/1206717593893740658/1246188695326035970/1249188290016579624"
             class="text-cl6 hover:underline">
          https://discord.com/channels/1206717593893740658/1246188695326035970/1249188290016579624
        </a>
        <br/>
        If you are unable to access this link, make sure you are in the EGTS Discord server. You can get the invite link at the bottom of this guide.
        <br/>
        <strong>
          If you are downloading for the first time, use the full download. The quick download is only for
          updating outdated versions!
        </strong>
      </p>
      <p class="mb-4">
        <strong>
          If you don&#39;t want to download the EGTS PACK or you don&#39;t want to connect to EGTS, you can setup your
          own local server using TLS.
          <br/>
          Instructions on how to set that up are in the EXTRA section where it says "TLS SETUP"
          in the Nijiiro setup guide.
        </strong>
      </p>
      <p class="mb-4">
        <strong class="underline">For support with EGTS, or questions and suggestions, join the EGTS discord server:</strong>
        <br/>
        <strong>
          <a href="https://discord.gg/egts" class="text-cl6 hover:underline">https://discord.gg/egts</a>
        </strong>
      </p>
      <p class="mb-4">
        <strong class="underline">Installing EGTS for the first time:</strong>
        <br/>
        Extract EGTSFullDownload.zip into the folder where you extracted 39.06. (E.G. C:\TaikoNijiiro)
        <br/>
        When asked to replace files, make sure you choose <strong>"REPLACE FILES"</strong>.
        <br/>
        If you did not get asked to replace files you did something wrong.
      </p>
      <p class="mb-4">
        <strong class="underline">Updating EGTS to a newer version:</strong>
        <br/>
        Follow the same procedures as if you were installing EGTS for the first time, this time using the Quick
        Download.
        <br/>
        If the Quick Download is not available or if your version is too old to support the quick download (if the
        date of your current version is not between the dates of the updates it contains), use the Full Download instead.
        <br/>
        You will be able to tell if your current version is outdated if the game tells you that you are using an
        outdated client in the song select.
      </p>
      <p class="mb-4">
        Open config.toml located in Executable\ with notepad and change <strong>server = "127.0.0.1"</strong>
        to <strong>server = "<a class="text-cl6 hover:underline" href="https://egts.ca/">egts.ca</a>"</strong>
        <br/>
        <br/>
        <strong>TAIKO MIGHT TAKE EXTRA LONG TO LOAD WHEN LOADING FOR THE FIRST TIME. (~90s)</strong>
      </p>
      <p class="mb-4">
        After you play 1 credit, you can register your card in the WebUI.
        <br/>
        EGTS WebUI:
        <br/>
        <a href="https://egts.ca/" class="text-cl6 hover:underline">https://egts.ca/</a>
      </p>
      <p class="mb-4">
        EGTS Discord:
        <br/>
        <a href="https://discord.gg/egts" class="text-cl6 hover:underline">https://discord.gg/egts</a>
      </p>
      <p>
        Taiko no Tatsujin Modding Discord Server:
        <br/>
        <a href="https://discord.gg/HFm37aA5zr" class="text-cl6 hover:underline">https://discord.gg/HFm37aA5zr</a>
      </p>
    </div>
</template>

<style scoped>

</style>