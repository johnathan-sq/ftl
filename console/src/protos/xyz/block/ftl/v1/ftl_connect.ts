// @generated by protoc-gen-connect-es v0.9.1 with parameter "target=ts"
// @generated from file xyz/block/ftl/v1/ftl.proto (package xyz.block.ftl.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CallRequest, CallResponse, CreateDeploymentRequest, CreateDeploymentResponse, DeployRequest, DeployResponse, GetArtefactDiffsRequest, GetArtefactDiffsResponse, GetDeploymentArtefactsRequest, GetDeploymentArtefactsResponse, GetDeploymentRequest, GetDeploymentResponse, PingRequest, PingResponse, RegisterRunnerRequest, RegisterRunnerResponse, SendMetricRequest, SendMetricResponse, StartDeployRequest, StartDeployResponse, StreamDeploymentLogsRequest, StreamDeploymentLogsResponse, TerminateRequest, TerminateResponse, UploadArtefactRequest, UploadArtefactResponse } from "./ftl_pb.js";
import { MethodIdempotency, MethodKind } from "@bufbuild/protobuf";

/**
 * VerbService is a common interface shared by multiple services for calling Verbs.
 *
 * @generated from service xyz.block.ftl.v1.VerbService
 */
export const VerbService = {
  typeName: "xyz.block.ftl.v1.VerbService",
  methods: {
    /**
     * Ping service for readiness.
     *
     * @generated from rpc xyz.block.ftl.v1.VerbService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * Issue a synchronous call to a Verb.
     *
     * @generated from rpc xyz.block.ftl.v1.VerbService.Call
     */
    call: {
      name: "Call",
      I: CallRequest,
      O: CallResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * @generated from service xyz.block.ftl.v1.ControlPlaneService
 */
export const ControlPlaneService = {
  typeName: "xyz.block.ftl.v1.ControlPlaneService",
  methods: {
    /**
     * Ping service for readiness.
     *
     * @generated from rpc xyz.block.ftl.v1.ControlPlaneService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * Get list of artefacts that differ between the server and client.
     *
     * @generated from rpc xyz.block.ftl.v1.ControlPlaneService.GetArtefactDiffs
     */
    getArtefactDiffs: {
      name: "GetArtefactDiffs",
      I: GetArtefactDiffsRequest,
      O: GetArtefactDiffsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Upload an artefact to the server.
     *
     * @generated from rpc xyz.block.ftl.v1.ControlPlaneService.UploadArtefact
     */
    uploadArtefact: {
      name: "UploadArtefact",
      I: UploadArtefactRequest,
      O: UploadArtefactResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Create a deployment.
     *
     * @generated from rpc xyz.block.ftl.v1.ControlPlaneService.CreateDeployment
     */
    createDeployment: {
      name: "CreateDeployment",
      I: CreateDeploymentRequest,
      O: CreateDeploymentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Get the schema and artefact metadata for a deployment.
     *
     * @generated from rpc xyz.block.ftl.v1.ControlPlaneService.GetDeployment
     */
    getDeployment: {
      name: "GetDeployment",
      I: GetDeploymentRequest,
      O: GetDeploymentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Stream deployment artefacts from the server.
     *
     * Each artefact is streamed one after the other as a sequence of max 1MB
     * chunks.
     *
     * @generated from rpc xyz.block.ftl.v1.ControlPlaneService.GetDeploymentArtefacts
     */
    getDeploymentArtefacts: {
      name: "GetDeploymentArtefacts",
      I: GetDeploymentArtefactsRequest,
      O: GetDeploymentArtefactsResponse,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * Register a Runner with the ControlPlane.
     *
     * Each runner issue a RegisterRunnerRequest to the ControlPlaneService
     * every 10 seconds to maintain its heartbeat.
     *
     * @generated from rpc xyz.block.ftl.v1.ControlPlaneService.RegisterRunner
     */
    registerRunner: {
      name: "RegisterRunner",
      I: RegisterRunnerRequest,
      O: RegisterRunnerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Starts a deployment.
     *
     * @generated from rpc xyz.block.ftl.v1.ControlPlaneService.StartDeploy
     */
    startDeploy: {
      name: "StartDeploy",
      I: StartDeployRequest,
      O: StartDeployResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Stream logs from a deployment
     *
     * @generated from rpc xyz.block.ftl.v1.ControlPlaneService.StreamDeploymentLogs
     */
    streamDeploymentLogs: {
      name: "StreamDeploymentLogs",
      I: StreamDeploymentLogsRequest,
      O: StreamDeploymentLogsResponse,
      kind: MethodKind.ClientStreaming,
    },
  }
} as const;

/**
 * RunnerService is the service that executes Deployments.
 *
 * The ControlPlane will scale the Runner horizontally as required. The Runner will
 * register itself automatically with the ControlPlaneService, which will then
 * assign modules to it.
 *
 * @generated from service xyz.block.ftl.v1.RunnerService
 */
export const RunnerService = {
  typeName: "xyz.block.ftl.v1.RunnerService",
  methods: {
    /**
     * @generated from rpc xyz.block.ftl.v1.RunnerService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * Initiate a deployment on this Runner.
     *
     * @generated from rpc xyz.block.ftl.v1.RunnerService.Deploy
     */
    deploy: {
      name: "Deploy",
      I: DeployRequest,
      O: DeployResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Terminate the deployment on this Runner.
     *
     * @generated from rpc xyz.block.ftl.v1.RunnerService.Terminate
     */
    terminate: {
      name: "Terminate",
      I: TerminateRequest,
      O: TerminateResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * @generated from service xyz.block.ftl.v1.ObservabilityService
 */
export const ObservabilityService = {
  typeName: "xyz.block.ftl.v1.ObservabilityService",
  methods: {
    /**
     * Ping service for readiness.
     *
     * @generated from rpc xyz.block.ftl.v1.ObservabilityService.Ping
     */
    ping: {
      name: "Ping",
      I: PingRequest,
      O: PingResponse,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc xyz.block.ftl.v1.ObservabilityService.SendMetric
     */
    sendMetric: {
      name: "SendMetric",
      I: SendMetricRequest,
      O: SendMetricResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

