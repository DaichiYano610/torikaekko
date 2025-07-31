import React, { useState } from 'react';
import { View, StyleSheet, Image, Alert } from 'react-native';
import { Input, Button, Card, Text, Icon } from 'react-native-elements';

export default function ItemExchange() {
  const [username, setUsername] = useState('');
  const [itemName, setItemName] = useState('');
  const [itemWant, setItemWant] = useState('');
  const [imageUri, setImageUri] = useState(null);

  const onImageSelect = () => {
    // 本来はImagePickerなど使う。今回はモックとしてURI設定
    setImageUri('https://via.placeholder.com/150');
  };

  const onSubmit = () => {
    Alert.alert(
      '送信内容',
      `ユーザーネーム: ${username}\n商品名: ${itemName}\n欲しい物: ${itemWant}`
    );
  };

  return (
    <View style={styles.container}>
      <Card containerStyle={styles.card}>
        <Card.Title>アイテム交換ポスト</Card.Title>
        <Card.Divider />

        <View style={styles.imageContainer}>
          {imageUri ? (
            <Image source={{ uri: imageUri }} style={styles.image} />
          ) : (
            <Button
              title="画像を選択"
              type="outline"
              onPress={onImageSelect}
              icon={<Icon name="image" type="feather" color="#2089dc" />}
            />
          )}
        </View>

        <Input
          label="商品名"
          placeholder="例: 自転車"
          value={itemName}
          onChangeText={setItemName}
        />
        <Input
          label="ほしい物"
          placeholder="例: ギター"
          value={itemWant}
          onChangeText={setItemWant}
        />
        <Input
          label="ユーザーネーム"
          placeholder="ニックネームを入力"
          value={username}
          onChangeText={setUsername}
        />

        <Button title="送信" onPress={onSubmit} containerStyle={styles.button} />
      </Card>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    backgroundColor: '#f5f5f5',
  },
  card: {
    borderRadius: 10,
    padding: 20,
  },
  imageContainer: {
    alignItems: 'center',
    marginVertical: 15,
  },
  image: {
    width: 150,
    height: 150,
    borderRadius: 8,
  },
  button: {
    marginTop: 20,
  },
});
