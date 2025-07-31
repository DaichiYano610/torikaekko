import React from 'react';
import { View, FlatList, TouchableOpacity } from 'react-native';
import { Card, Text } from 'react-native-elements';
import { useRouter } from 'expo-router';

const items = [
  { id: '1', name: '自転車', description: '街乗りに最適な自転車' },
  { id: '2', name: 'ギター', description: 'アコースティックギター（初心者向け）' },
  { id: '3', name: 'ノートパソコン', description: '軽量で持ち運び便利なPC' },
];

export default function ItemListScreen() {
  const router = useRouter();

  const goToDetail = (id: string) => {
    router.push(`/items/${id}`);
  };

  return (
    <View style={{ flex: 1, padding: 16 }}>
      <FlatList
        data={items}
        keyExtractor={(item) => item.id}
        renderItem={({ item }) => (
          <TouchableOpacity onPress={() => goToDetail(item.id)}>
            <Card containerStyle={{ borderRadius: 10 }}>
              <Card.Title>{item.name}</Card.Title>
              <Card.Divider />
              <Text style={{ fontSize: 14, color: '#555' }}>{item.description}</Text>
            </Card>
          </TouchableOpacity>
        )}
      />
    </View>
  );
}
