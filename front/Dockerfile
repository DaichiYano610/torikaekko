FROM node:lts-alpine
# 作業ディレクトリ作成
WORKDIR /usr/src/app/

# ビルドに必要なパッケージを追加
RUN apk update && apk add --no-cache bash python3 make g++ git


# アプリのパッケージ定義をコピー
COPY package.json yarn.lock ./
# Expo SDK を含む依存パッケージをインストール
RUN yarn install

# ソースコードをコピー
COPY . .



#for test code 
CMD [ "bash" ]

# Expoプロジェクトを起動
# CMD ["yarn", "expo", "start", "--tunnel"]

