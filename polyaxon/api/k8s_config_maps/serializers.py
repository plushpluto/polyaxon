from api.utils.serializers.catalog import K8SResourceCatalogSerializer, CatalogNameSerializer
from db.models.config_maps import K8SConfigMap


class K8SConfigMapSerializer(K8SResourceCatalogSerializer):
    QUERY = K8SConfigMap.objects

    class Meta:
        model = K8SConfigMap
        fields = K8SResourceCatalogSerializer.Meta.fields


class K8SConfigMapNameSerializer(CatalogNameSerializer):
    class Meta:
        model = K8SConfigMap
        fields = CatalogNameSerializer.Meta.fields
