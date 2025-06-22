import React, { useState } from 'react';
import { View, StyleSheet } from 'react-native';
import { Input, Button } from 'react-native-elements';

export default function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const onSubmit = () => {
    alert(`ユーザーネーム: ${username}\nパスワード: ${password}`);
  };

  return (
    <View >
      <Input
        label="ユーザーネーム"
        placeholder="ユーザーネームを入力"
        value={username}
        onChangeText={setUsername}
        autoCapitalize="none"
      />
      <Input
        label="パスワード"
        placeholder="パスワードを入力"
        value={password}
        onChangeText={setPassword}
        secureTextEntry
        autoCapitalize="none"
      />
      <Button title="ログイン" onPress={onSubmit} />
    </View>
  );
}



