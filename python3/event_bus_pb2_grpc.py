# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from . import event_bus_pb2 as event__bus__pb2
from . import miscellaneous_pb2 as miscellaneous__pb2


class EventBusStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.JoinEventBus = channel.unary_unary(
        '/dialog.EventBus/JoinEventBus',
        request_serializer=event__bus__pb2.RequestJoinEventBus.SerializeToString,
        response_deserializer=event__bus__pb2.ResponseJoinEventBus.FromString,
        )
    self.KeepAliveEventBus = channel.unary_unary(
        '/dialog.EventBus/KeepAliveEventBus',
        request_serializer=event__bus__pb2.RequestKeepAliveEventBus.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )
    self.PostToEventBus = channel.unary_unary(
        '/dialog.EventBus/PostToEventBus',
        request_serializer=event__bus__pb2.RequestPostToEventBus.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )


class EventBusServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def JoinEventBus(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def KeepAliveEventBus(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PostToEventBus(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_EventBusServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'JoinEventBus': grpc.unary_unary_rpc_method_handler(
          servicer.JoinEventBus,
          request_deserializer=event__bus__pb2.RequestJoinEventBus.FromString,
          response_serializer=event__bus__pb2.ResponseJoinEventBus.SerializeToString,
      ),
      'KeepAliveEventBus': grpc.unary_unary_rpc_method_handler(
          servicer.KeepAliveEventBus,
          request_deserializer=event__bus__pb2.RequestKeepAliveEventBus.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
      'PostToEventBus': grpc.unary_unary_rpc_method_handler(
          servicer.PostToEventBus,
          request_deserializer=event__bus__pb2.RequestPostToEventBus.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'dialog.EventBus', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
