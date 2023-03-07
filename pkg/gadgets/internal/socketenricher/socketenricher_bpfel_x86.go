// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package socketenricher

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type socketenricherSocketsKey struct {
	Netns  uint32
	Family uint16
	Proto  uint16
	Port   uint16
	_      [2]byte
}

type socketenricherSocketsValue struct {
	Mntns   uint64
	PidTgid uint64
	Task    [16]int8
}

// loadSocketenricher returns the embedded CollectionSpec for socketenricher.
func loadSocketenricher() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_SocketenricherBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load socketenricher: %w", err)
	}

	return spec, err
}

// loadSocketenricherObjects loads socketenricher and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*socketenricherObjects
//	*socketenricherPrograms
//	*socketenricherMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadSocketenricherObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadSocketenricher()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// socketenricherSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type socketenricherSpecs struct {
	socketenricherProgramSpecs
	socketenricherMapSpecs
}

// socketenricherSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type socketenricherProgramSpecs struct {
	IgBindIpv4E   *ebpf.ProgramSpec `ebpf:"ig_bind_ipv4_e"`
	IgBindIpv4X   *ebpf.ProgramSpec `ebpf:"ig_bind_ipv4_x"`
	IgBindIpv6E   *ebpf.ProgramSpec `ebpf:"ig_bind_ipv6_e"`
	IgBindIpv6X   *ebpf.ProgramSpec `ebpf:"ig_bind_ipv6_x"`
	IgFreeIpv4E   *ebpf.ProgramSpec `ebpf:"ig_free_ipv4_e"`
	IgFreeIpv6E   *ebpf.ProgramSpec `ebpf:"ig_free_ipv6_e"`
	IgTcpcV4CoE   *ebpf.ProgramSpec `ebpf:"ig_tcpc_v4_co_e"`
	IgTcpcV4CoX   *ebpf.ProgramSpec `ebpf:"ig_tcpc_v4_co_x"`
	IgTcpcV6CoE   *ebpf.ProgramSpec `ebpf:"ig_tcpc_v6_co_e"`
	IgTcpcV6CoX   *ebpf.ProgramSpec `ebpf:"ig_tcpc_v6_co_x"`
	IgUdp6Sendmsg *ebpf.ProgramSpec `ebpf:"ig_udp6_sendmsg"`
	IgUdpSendmsg  *ebpf.ProgramSpec `ebpf:"ig_udp_sendmsg"`
}

// socketenricherMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type socketenricherMapSpecs struct {
	Sockets *ebpf.MapSpec `ebpf:"sockets"`
	Start   *ebpf.MapSpec `ebpf:"start"`
}

// socketenricherObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadSocketenricherObjects or ebpf.CollectionSpec.LoadAndAssign.
type socketenricherObjects struct {
	socketenricherPrograms
	socketenricherMaps
}

func (o *socketenricherObjects) Close() error {
	return _SocketenricherClose(
		&o.socketenricherPrograms,
		&o.socketenricherMaps,
	)
}

// socketenricherMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadSocketenricherObjects or ebpf.CollectionSpec.LoadAndAssign.
type socketenricherMaps struct {
	Sockets *ebpf.Map `ebpf:"sockets"`
	Start   *ebpf.Map `ebpf:"start"`
}

func (m *socketenricherMaps) Close() error {
	return _SocketenricherClose(
		m.Sockets,
		m.Start,
	)
}

// socketenricherPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadSocketenricherObjects or ebpf.CollectionSpec.LoadAndAssign.
type socketenricherPrograms struct {
	IgBindIpv4E   *ebpf.Program `ebpf:"ig_bind_ipv4_e"`
	IgBindIpv4X   *ebpf.Program `ebpf:"ig_bind_ipv4_x"`
	IgBindIpv6E   *ebpf.Program `ebpf:"ig_bind_ipv6_e"`
	IgBindIpv6X   *ebpf.Program `ebpf:"ig_bind_ipv6_x"`
	IgFreeIpv4E   *ebpf.Program `ebpf:"ig_free_ipv4_e"`
	IgFreeIpv6E   *ebpf.Program `ebpf:"ig_free_ipv6_e"`
	IgTcpcV4CoE   *ebpf.Program `ebpf:"ig_tcpc_v4_co_e"`
	IgTcpcV4CoX   *ebpf.Program `ebpf:"ig_tcpc_v4_co_x"`
	IgTcpcV6CoE   *ebpf.Program `ebpf:"ig_tcpc_v6_co_e"`
	IgTcpcV6CoX   *ebpf.Program `ebpf:"ig_tcpc_v6_co_x"`
	IgUdp6Sendmsg *ebpf.Program `ebpf:"ig_udp6_sendmsg"`
	IgUdpSendmsg  *ebpf.Program `ebpf:"ig_udp_sendmsg"`
}

func (p *socketenricherPrograms) Close() error {
	return _SocketenricherClose(
		p.IgBindIpv4E,
		p.IgBindIpv4X,
		p.IgBindIpv6E,
		p.IgBindIpv6X,
		p.IgFreeIpv4E,
		p.IgFreeIpv6E,
		p.IgTcpcV4CoE,
		p.IgTcpcV4CoX,
		p.IgTcpcV6CoE,
		p.IgTcpcV6CoX,
		p.IgUdp6Sendmsg,
		p.IgUdpSendmsg,
	)
}

func _SocketenricherClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed socketenricher_bpfel_x86.o
var _SocketenricherBytes []byte
