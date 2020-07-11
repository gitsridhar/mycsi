package mycsi

import (
        "path/filepath"
        "fmt"
        "os"
        "sync"
        "strings"
        "time"

        "github.com/golang/glog"
        "golang.org/x/net/context"

        "github.com/container-storage-interface/spec/lib/go/csi"
        "github.com/nightlyone/lockfile"
        "google.golang.org/grpc/codes"
        "google.golang.org/grpc/status"
)

type NodeServer struct {
        Driver  *MyCSIDriver
        // TODO: Only lock mutually exclusive calls and make locking more fine grained
        mux sync.Mutex
}

func (ns *NodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
        glog.V(3).Infof("nodeserver NodePublishVolume")
        glog.V(4).Infof("NodePublishVolume called with req: %#v", req)

        // Validate Arguments
        volumeID := req.GetVolumeId()
        targetPath := req.GetTargetPath()
        stagingTargetPath := req.GetStagingTargetPath()
        volumeCapability := req.GetVolumeCapability()
        if len(volumeID) == 0 {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume ID must be provided")
        }
        if len(targetPath) == 0 {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Target Path must be provided")
        }
        if len(stagingTargetPath) == 0 {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Staging Target Path must be provided")
        }
        if volumeCapability == nil {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume Capability must be provided")
        }

        return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *NodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
        glog.V(3).Infof("nodeserver NodeUnpublishVolume")
        glog.V(4).Infof("NodeUnpublishVolume called with args: %v", req)

        // Validate Arguments
        volumeID := req.GetVolumeId()
        targetPath := req.GetTargetPath()
        if len(volumeID) == 0 {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume ID must be provided")
        }
        if len(targetPath) == 0 {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Staging Target Path must be provided")
        }

        return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *NodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (
        *csi.NodeStageVolumeResponse, error) {
        glog.V(3).Infof("nodeserver NodeStageVolume %#v", req)

        // Validate Arguments
        volumeID := req.GetVolumeId()
        stagingTargetPath := req.GetStagingTargetPath()
        volumeCapability := req.GetVolumeCapability()
        if len(volumeID) == 0 {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume ID must be provided")
        }
        if len(stagingTargetPath) == 0 {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Staging Target Path must be provided")
        }
        if volumeCapability == nil {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume Capability must be provided")
        }

        fsType := volumeCapability.GetMount().FsType
        glog.V(3).Infof("nodeserver NodeStageVolume Required Filesystem Type : %s", fsType)

        var pID = os.Getpid()
        lock, err := lockfile.New(filepath.Join(os.TempDir(), resources.ScsiScanLock))
        if err != nil {
                return nil, status.Error(codes.InvalidArgument, (fmt.Sprintf("%d : Cannot init lock. Reason. %v", pID, err)))
        }
        
        for i := 0; i < resources.MaxAttemptsToTryLock; i++ {
                err = lock.TryLock()
                if err == nil {
                        break
                }
                glog.V(3).Infof("%d : Could not get lock, error is %v. Sleeping for 5 seconds", pID, err)
                time.Sleep( 5 * time.Second)
        }
        glog.V(3).Infof("%d : Got hold of Scsiscan lock", pID)
        
        defer lock.Unlock()

        return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *NodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (
        *csi.NodeUnstageVolumeResponse, error) {
        glog.V(3).Infof("nodeserver NodeUnstageVolume")

        // Validate Arguments
        volumeID := req.GetVolumeId()
        stagingTargetPath := req.GetStagingTargetPath()
        if len(volumeID) == 0 {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Volume ID must be provided")
        }
        if len(stagingTargetPath) == 0 {
                return nil, status.Error(codes.InvalidArgument, "NodeStageVolume Staging Target Path must be provided")
        }

        return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *NodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
        glog.V(4).Infof("NodeGetCapabilities called with req: %#v", req)
        return &csi.NodeGetCapabilitiesResponse{
                Capabilities: ns.Driver.nscap,
        }, nil
}

func (ns *NodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
        glog.V(4).Infof("NodeGetInfo called with req: %#v", req)
        return &csi.NodeGetInfoResponse{
                NodeId: ns.Driver.nodeID,
        }, nil
}

func (ns *NodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {

        volumeID := req.GetVolumeId()
        if volumeID == "" {
                return nil, status.Error(codes.InvalidArgument, "VolumeID is not present")
        }

        newSize := int(req.GetCapacityRange().GetRequiredBytes() / GiB)

        return &csi.NodeExpandVolumeResponse{
        }, nil

}
func (ns *NodeServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
        return nil, status.Error(codes.Unimplemented, "")
}
