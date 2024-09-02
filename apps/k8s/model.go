package k8s

func (p *GetPodRamUsageRsp) Len() int {
	return int(p.Count)
}
func (p *GetPodRamUsageRsp) Less(i, j int) bool {
	return p.Items[i].RamMbi > p.Items[j].RamMbi
}
func (p *GetPodRamUsageRsp) Swap(i, j int) {
	p.Items[i], p.Items[j] = p.Items[j], p.Items[i]
}
