#include "authz.h"

#include <google/protobuf/util/json_util.h>
#include <google/rpc/code.pb.h>

#include "err/errors.h"

namespace svc {
namespace authz {
template <>
rpcGrant::result_type Impl::call<rpcGrant>(
	grpcxx::context &ctx, const rpcGrant::request_type &req) {
	// Upsert if exists
	if (auto r = db::LookupRecord(req.principal_id(), req.resource_type(), req.resource_id()); r) {
		if (req.has_attrs()) {
			std::string attrs;
			google::protobuf::util::MessageToJsonString(req.attrs(), &attrs);

			r->attrs(std::move(attrs));
			r->store();
		}

		return {grpcxx::status::code_t::ok, map(r.value())};
	}

	auto r = map(req);
	r.store();

	return {grpcxx::status::code_t::ok, map(r)};
}

google::rpc::Status Impl::exception() noexcept {
	google::rpc::Status status;
	status.set_code(google::rpc::UNKNOWN);

	try {
		std::rethrow_exception(std::current_exception());
	} catch (const err::DbRecordInvalidPrincipalId &e) {
		status.set_code(google::rpc::INVALID_ARGUMENT);
		status.set_message(std::string(e.str()));
	} catch (const err::DbRevisionMismatch &e) {
		status.set_code(google::rpc::INTERNAL);
		status.set_message(std::string(e.str()));
	} catch (const std::exception &e) {
		status.set_code(google::rpc::INTERNAL);
		status.set_message(e.what());
	}

	return status;
}

db::Record Impl::map(const rpcGrant::request_type &from) const noexcept {
	db::Record to({
		.principalId  = from.principal_id(),
		.resourceId   = from.resource_id(),
		.resourceType = from.resource_type(),
	});

	if (from.has_attrs()) {
		std::string attrs;
		google::protobuf::util::MessageToJsonString(from.attrs(), &attrs);

		to.attrs(std::move(attrs));
	}

	return to;
}

rpcGrant::response_type Impl::map(const db::Record &from) const noexcept {
	return {};
}
} // namespace authz
} // namespace svc
