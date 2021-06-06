### view
- カード：属性（表裏状態、カード識別）
- カードはカード識別が各２セット
- カード識別の枚数は可変。ランダムに配置される。

### game
- プレイヤー数は２人以上。はじめに名前入力してもらう。
- プレイヤーは自分のターンに裏向きのカードを２枚選択する。選択したカードが同じ場合は自分の持ち札になり、再度選択できる。選択後、次のプレイヤーターンとなる。
- 持ち札となったカードは盤面から自分の手札に移る。盤面上のスペースはそのまま。（詰めない）
- 選択したカードが持ち札となり、そのあと選択可能なカードが無い場合はゲームセットとなる
- ゲームセット時に各プレイヤーの持ち札のペア数を集計し、もっとも多いプレイヤーが勝利。
- ゲームセット時に各プレイヤーの持ち札が整理されて表示されるとCool！！

### data
- プレイヤーをオブジェクト化してデータや状態を保持させる
    - 持ち札
    - プレイヤーをローテートすることでターンを表現する
- メタデータ
    - 経過ターン数