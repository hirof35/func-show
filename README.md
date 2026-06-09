# Trigonometric Function Simulator (func-show)

三角関数（$\sin$, $\cos$）の挙動を単位円とスライダを用いて直感的に理解するための、マルチプラットフォーム対応の視覚化シミュレーターです。
<img width="492" height="591" alt="スクリーンショット 2026-06-10 072227" src="https://github.com/user-attachments/assets/09e0671e-c629-499b-b71b-94361de6350c" />

オリジナルの Python (PyGame) 実装から、Go, Swift, JavaScript の各環境へ移植・最適化されています。

## 💡 概要

スライダを動かして角度（0° 〜 360°）を変更することで、単位円上の点における以下の要素がリアルタイムに連動・描画されます。
* **$\cos$ (青線)**: X軸への射影（垂線）
* **$\sin$ (緑線)**: Y軸への射影（垂線）
* **半径 (赤線)**: 中心から円周上の点へのベクトル

すべての環境において、画面の座標系（左上原点）から数学的な座標系（第一象限が右上）への補正が施されており、滑らかな操作感を実現しています。

---

## 🛠️ 各言語の実行方法

### 1. Go (Ebitengine)
2Dゲームエンジン `Ebitengine` を使用したネイティブデスクトップアプリ実装です。

#### 依存パッケージの導入
```bash
go mod init func-show
go get [github.com/hajimehoshi/ebiten/v2](https://github.com/hajimehoshi/ebiten/v2)
実行方法
Bash
go run main.go
2. JavaScript (HTML5 Canvas)
ブラウザ上で環境構築なしに即座に動作する、Vanilla JS 実装です。

実行方法
index.html を作成し、ソースコードを貼り付けます。

ファイルをダブルクリック、またはブラウザにドラッグ＆ドロップして開きます。

3. Swift (SwiftUI)
Appleプラットフォーム向けに状態管理（State）を最適化した、モダンな宣言型UI実装です。

実行方法
Xcode を起動し、新規の SwiftUI Project を作成します。

ContentView.swift にソースコードを配置し、プレビューまたはシミュレーターで実行します。

📐 構造論的特徴（リファクタリングの要点）
座標系の一貫性:
グラフィックスライブラリ特有の「下方向がY軸の正」という特性を考慮し、すべての移植先で ypos = CenterY - (sin(rad) * Radius) による反転処理を統一して組み込んでいます。

イベント駆動・リアクティブ化:
PyGame版の固定フレームレート（10 FPS）から、JSの requestAnimationFrame や SwiftUIのデータバインディングへと移行し、描画パフォーマンスと応答性を大幅に向上させています。
