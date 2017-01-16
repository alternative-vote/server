package generated

type SerializationMetadata struct {
	deserializedProperty []string
  fieldsToSerialize []string
}

func (o *SerializationMetadata) AddDeserializedProperty(s string) error {
	o.deserializedProperty = append(o.deserializedProperty, s)
	return nil
}

func (o *SerializationMetadata) GetDeserializedProperties() []string {
	return o.deserializedProperty
}

func (o *SerializationMetadata) SetFieldsToSerialize(ss []string) {
	o.fieldsToSerialize = ss
}

