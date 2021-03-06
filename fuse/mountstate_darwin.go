package fuse

func (ms *MountState) systemWrite(req *request, header []byte) Status {
	if req.flatDataSize() == 0 {
		_, err := ms.mountFile.Write(header)
		return ToStatus(err)
	}

	if req.fdData != nil {
		sz := req.flatDataSize()
		buf := ms.AllocOut(req, uint32(sz))
		req.flatData, req.status = req.fdData.Bytes(buf)
		header = req.serializeHeader(len(req.flatData))
	}

	_, err := Writev(int(ms.mountFile.Fd()), [][]byte{header, req.flatData})
	if req.readResult != nil {
		req.readResult.Done()
	}
	return ToStatus(err)
}
