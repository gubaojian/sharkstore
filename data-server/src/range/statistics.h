_Pragma("once");

#include "monitor/syscommon.h"

namespace sharkstore {
namespace dataserver {
namespace range {

class RangeStats {
public:
    RangeStats() = default;
    virtual ~RangeStats() = default;

    virtual void PushTime(monitor::PrintTag type, uint32_t time) {}
    virtual void IncrLeaderCount() {}
    virtual void DecrLeaderCount() {}
    virtual void IncrSplitCount() {}
    virtual void DecrSplitCount() {}

    virtual uint64_t GetFilesystemUsedPercent() const { return 0; }
};

}  // namespace range
}  // namespace dataserver
}  // namespace sharkstore
