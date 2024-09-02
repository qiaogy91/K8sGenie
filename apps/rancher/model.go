package rancher

func (ps *ProjectSet) ProjectIds() []string {
	var ids []string
	for _, pro := range ps.Items {
		ids = append(ids, pro.Spec.ProjectId)
	}
	return ids
}
