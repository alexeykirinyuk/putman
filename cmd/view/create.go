package view

import "github.com/alexeykirinyuk/putman/application"

func CreateViews(
	collectionService *application.CollectionService,
) []IView {
	return []IView{
		TreeView{
			collectionService: collectionService,
		},
		AddCollectionView{
			collectionService: collectionService,
		},
	}
}
