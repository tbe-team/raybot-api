# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from raybot.v1 import auth_pb2 as raybot_dot_v1_dot_auth__pb2


class AuthServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetSession = channel.unary_unary(
                '/raybot.v1.AuthService/GetSession',
                request_serializer=raybot_dot_v1_dot_auth__pb2.GetSessionRequest.SerializeToString,
                response_deserializer=raybot_dot_v1_dot_auth__pb2.GetSessionResponse.FromString,
                _registered_method=True)


class AuthServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetSession(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AuthServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetSession': grpc.unary_unary_rpc_method_handler(
                    servicer.GetSession,
                    request_deserializer=raybot_dot_v1_dot_auth__pb2.GetSessionRequest.FromString,
                    response_serializer=raybot_dot_v1_dot_auth__pb2.GetSessionResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'raybot.v1.AuthService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('raybot.v1.AuthService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class AuthService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetSession(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/raybot.v1.AuthService/GetSession',
            raybot_dot_v1_dot_auth__pb2.GetSessionRequest.SerializeToString,
            raybot_dot_v1_dot_auth__pb2.GetSessionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
