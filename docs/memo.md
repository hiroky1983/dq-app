# Doragon Questのバトルを設計する

## 用意するもの
- コンテキスト図
- ユースケース図
- シーケンス図
- クラス図
- DB設計
- ユビキタス言語
- ドメインモデル


## 必要そうなドメイン
- 味方キャラ
- 味方パーティグループ
- モンスター
- 攻撃呪文
- 回復呪文

## 勇者の仕様
- 勇者は名前、HP、MP、攻撃力、防御力、すばやさ、会心率を持つ
- 勇者はすべてのステータスのバランスがよい

## キャラクターの仕様
- キャラクターは名前、HP、MP、攻撃力、防御力、すばやさ、会心率を持つ
- キャラクターは戦士、魔法使い、僧侶、武闘家のいずれか
- 戦士はHPが高く、攻撃力が高い、MPが0で呪文を使えない
- 魔法使いはMPが高く、攻撃呪文を使える、攻撃力、HPが低い
- 僧侶はHPが高く、MPが高い、回復呪文を使える
- 武闘家はすばやさが高く、攻撃力が高く、会心率が高い、MPが0で呪文を使えない

## パーティの仕様
- パーティは最大4人
- パーティは最大1人の勇者を持つ
- 残りの仲間は戦士、魔法使い、僧侶、武闘家のいずれか
- パーティは最大1人の勇者を持つ
- パーティの仲間の職業の重複は可
- 一番最初にパーティメンバーの選択をする

## モンスターの仕様
- モンスターは名前、HP、MP、攻撃力、防御力、すばやさを持つ
- モンスターの強さはグレードで表現する
- モンスターは1体のみ出現して、複数では出現しない
- モンスターはHPが0になると倒れる

## 攻撃呪文の仕様
- 属性は今回なしで、単純にダメージを与える

## 回復呪文の仕様
- 最大HPの40%~50%を回復する

## 職業の仕様
- 戦士、魔法使い、僧侶、武闘家から選択する
- 名前は自由につけられる
- 各職業によってステータスのテーブルが異なる

## バトルの仕様
1. モンスターが出現
2. プレイヤーは味方の行動を選択
3. プレイヤーとモンスターのすばやさを計算して、行動順を計算
4. 各行動を実行
5. ターンエンド
6. 敵のHPが0になるか、味方全員のHPが0になるとバトル終了

## やらないこと
- レベルアップ
- アイテム
- 武器、防具、装備
- お金
- マップ
- イベント
- ダンジョン
- エンカウント
- フィールド
- ショップ
- 転職
- スキル
- バフ、デバフ
- ステータス異常
