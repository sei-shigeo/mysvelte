<script>
  import * as App from "../wailsjs/go/main/App.js";
  import { fade } from "svelte/transition";

  let inputPDFPath = "";
  let stampImagePath = "";
  let outputDirectory = "";
  let outputFileName = "";
  let message = "";

  // 入力PDFを選択
  async function selectInputPDF() {
    inputPDFPath = await App.SelectFile();
  }

  // 印鑑画像を選択
  async function selectStampImage() {
    stampImagePath = await App.SelectFile();
  }

  // 出力先ディレクトリを選択
  async function selectOutputDirectory() {
    outputDirectory = await App.SelectDirectory();
  }

  // 印鑑を追加
  async function handleAddStamp() {
    if (!inputPDFPath || !stampImagePath || !outputDirectory) {
      message = "すべての必要な情報を入力してください。";
      return;
    }

    const outputPDFPath = `${outputDirectory}/${outputFileName}`;

    try {
      await App.AddStampToPDF(stampImagePath, inputPDFPath, outputPDFPath);
      message = "PDFに印鑑が正常に追加されました。";
    } catch (error) {
      message = `エラーが発生しました: ${error}`;
    }
    clearMessageAfterDelay();
  }
  function clearMessageAfterDelay() {
    setTimeout(() => {
      message = "";
    }, 3000); // 3秒後にメッセージを消す
  }
</script>

<main>
  <h1>PDF印鑑追加ツール</h1>

  <!-- 入力PDF -->
  <div class="input-group">
    <label for="inputPDF">入力PDF:</label>
    <div class="input-with-button">
      <input
        type="text"
        id="inputPDF"
        bind:value={inputPDFPath}
        readonly
        placeholder="PDFファイルを選択してください"
      />
      <button on:click={selectInputPDF}>選択</button>
    </div>
  </div>

  <!-- 印鑑画像 -->
  <div class="input-group">
    <label for="stampImage">印鑑画像:</label>
    <div class="input-with-button">
      <input
        type="text"
        id="stampImage"
        bind:value={stampImagePath}
        readonly
        placeholder="印鑑画像を選択してください"
      />
      <button on:click={selectStampImage}>選択</button>
    </div>
  </div>

  <!-- 出力先ディレクトリ -->
  <div class="input-group">
    <label for="outputDir">出力先ディレクトリ:</label>
    <div class="input-with-button">
      <input
        type="text"
        id="outputDir"
        bind:value={outputDirectory}
        readonly
        placeholder="出力先ディレクトリを選択してください"
      />
      <button on:click={selectOutputDirectory}>選択</button>
    </div>
  </div>

  <!-- 出力ファイル名 -->
  <div class="input-group">
    <label for="outputFileName">出力ファイル名:</label>
    <input
      type="text"
      id="outputFileName"
      bind:value={outputFileName}
      placeholder="出力ファイル名を入力してください"
    />
  </div>

  <!-- 印鑑を追加ボタン -->
  <button class="primary-button" on:click={handleAddStamp}>印鑑を追加</button>

  <!-- メッセージ -->
  {#if message}
    <p class="message" transition:fade>{message}</p>
  {/if}
</main>

<style>
  :global(body) {
    font-family: Arial, sans-serif;
    /* background-color: #f0f0f0; */
    margin: 0;
    padding: 0;
  }

  main {
    max-width: 600px;
    margin: 2rem auto;
    padding: 2rem;
    background-color: white;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  }

  h1 {
    text-align: center;
    color: #333;
    margin-bottom: 2rem;
  }

  .input-group {
    margin-bottom: 1rem;
  }

  label {
    display: block;
    margin-bottom: 0.5rem;
    color: #555;
  }

  input[type="text"] {
    width: 97%;
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 1rem;
  }

  .input-with-button {
    display: flex;
  }

  .input-with-button input[type="text"] {
    flex-grow: 1;
    border-top-right-radius: 0;
    border-bottom-right-radius: 0;
  }

  button {
    padding: 0.5rem 1rem;
    background-color: #4caf50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 1rem;
    transition: background-color 0.3s;
  }

  .input-with-button button {
    border-top-left-radius: 0;
    border-bottom-left-radius: 0;
  }

  button:hover {
    background-color: #45a049;
  }

  .primary-button {
    display: block;
    width: 100%;
    margin-top: 1rem;
    padding: 0.75rem;
    background-color: #007bff;
  }

  .primary-button:hover {
    background-color: #0056b3;
  }

  .message {
    position: absolute;
    top: 30px;
    right: 30px;
    margin-top: 1rem;
    padding: 0.5rem;
    background-color: #e6f7ff;
    border: 1px solid #91d5ff;
    border-radius: 4px;
    color: #0050b3;
  }
</style>
