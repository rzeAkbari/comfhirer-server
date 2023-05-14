package domain

type Bundle struct {
	/** Resource Type Name (for serialization) */
	ResourceType string `json:"resourceType"`
	/**
	 * An entry in a bundle resource - will either contain a resource or information about a resource (transactions and history only).
	 */
	Entry []BundleEntry `json:"entry,omitempty"`
	/**
	 * Persistent identity generally only matters for batches of type Document, Message, and Collection. It would not normally be populated for search and history results and servers ignore Bundle.identifier when processing batches and transactions. For Documents  the .identifier SHALL be populated such that the .identifier is globally unique.
	 */
	Identifier *Identifier `json:"identifier,omitempty"`
	/**
	 * Both Bundle.link and Bundle.entry.link are defined to support providing additional context when Bundles are used (e.g. [HATEOAS](http//en.wikipedia.org/wiki/HATEOAS)).
	 * Bundle.entry.link corresponds to links found in the HTTP header if the resource in the entry was [read](http.html#read) directly.
	 * This specification defines some specific uses of Bundle.link for [searching](search.html#conformance) and [paging](http.html#paging), but no specific uses for Bundle.entry.link, and no defined function in a transaction - the meaning is implementation specific.
	 */
	Link []BundleLink `json:"link,omitempty"`
	/**
	 * The signature could be created by the "author" of the bundle or by the originating device.   Requirements around inclusion of a signature, verification of signatures and treatment of signed/non-signed bundles is implementation-environment specific.
	 */
	Signature *Signature `json:"signature,omitempty"`
	/**
	 * For many bundles, the timestamp is equal to .meta.lastUpdated, because they are not stored (e.g. search results). When a bundle is placed in a persistent store, .meta.lastUpdated will be usually be changed by the server. When the bundle is a message, a middleware agent altering the message (even if not stored) SHOULD update .meta.lastUpdated. .timestamp is used to track the original time of the Bundle, and SHOULD be populated.
	 * Usage
	 * * document  the date the document was created. Note the composition may predate the document, or be associated with multiple documents. The date of the composition - the authoring time - may be earlier than the document assembly time
	 * * message  the date that the content of the message was assembled. This date is not changed by middleware engines unless they add additional data that changes the meaning of the time of the message
	 * * history  the date that the history was assembled. This time would be used as the _since time to ask for subsequent updates
	 * * searchset  the time that the search set was assembled. Note that different pages MAY have different timestamps but need not. Having different timestamps does not imply that subsequent pages will represent or include changes made since the initial query
	 * * transaction | transaction-response | batch | batch-response | collection  no particular assigned meaning
	 * The timestamp value should be greater than the lastUpdated and other timestamps in the resources in the bundle, and it should be equal or earlier than the .meta.lastUpdated on the Bundle itself.
	 */
	Timestamp     string   `json:"timestamp,omitempty"`
	Ext_Timestamp *Element `json:"_timestamp,omitempty"`
	/**
	 * Only used if the bundle is a search result set. The total does not include resources such as OperationOutcome and included resources, only the total number of matching resources.
	 */
	Total int64 `json:"total,omitempty"`
	/**
	 * It's possible to use a bundle for other purposes (e.g. a document can be accepted as a transaction). This is primarily defined so that there can be specific rules for some of the bundle types.
	 */
	Type     string   `json:"type,omitempty"`
	Ext_type *Element `json:"_type,omitempty"`
}

type EntryFhirResource interface {
}

type BundleEntry struct {
	/**
	 * fullUrl might not be [unique in the context of a resource](bundle.html#bundle-unique). Note that since [FHIR resources do not need to be served through the FHIR API](references.html), the fullURL might be a URN or an absolute URL that does not end with the logical id of the resource (Resource.id). However, but if the fullUrl does look like a RESTful server URL (e.g. meets the [regex](references.html#regex), then the 'id' portion of the fullUrl SHALL end with the Resource.id.
	 * Note that the fullUrl is not the same as the canonical URL - it's an absolute url for an endpoint serving the resource (these will happen to have the same value on the canonical server for the resource with the canonical URL).
	 */
	FullUrl     string   `json:"fullUrl,omitempty"`
	Ext_fullUrl *Element `json:"_fullUrl,omitempty"`
	/**
	 * A series of links that provide context to this entry.
	 */
	Link []BundleLink `json:"link,omitempty"`
	/**
	 * Additional information about how this entry should be processed as part of a transaction or batch.  For history, it shows how the entry was processed to create the version contained in the entry.
	 */
	Request *BundleEntryRequest `json:"request,omitempty"`
	/**
	 * The Resource for the entry. The purpose/meaning of the resource is determined by the Bundle.type.
	 */
	Resource EntryFhirResource `json:"resource"`
	/**
	 * Indicates the results of processing the corresponding 'request' entry in the batch or transaction being responded to or what the results of an operation where when returning history.
	 */
	Response *BundleEntryResponse `json:"response,omitempty"`
	/**
	 * Information about the search process that lead to the creation of this entry.
	 */
	Search *BundleEntrySearch `json:"search,omitempty"`
}

type BundleLink struct {

	/**
	 * A name which details the functional use for this link - see [http//www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1](http//www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1).
	 */
	Relation     string   `json:"relation,omitempty"`
	Ext_relation *Element `json:"_relation,omitempty"`
	/**
	 * The reference details for the link.
	 */
	Url     string   `json:"url,omitempty"`
	Ext_url *Element `json:"_url,omitempty"`
}

type BundleSearchMode string

const (
	Match   BundleSearchMode = "match"
	Include BundleSearchMode = "include"
	outcome BundleSearchMode = "outcome"
)

type BundleEntrySearch struct {
	/**
	 * There is only one mode. In some corner cases, a resource may be included because it is both a match and an include. In these circumstances, 'match' takes precedence.
	 */
	Mode     *BundleSearchMode `json:"mode,omitempty"`
	Ext_mode *Element          `json:"_mode,omitempty"`
	/**
	 * Servers are not required to return a ranking score. 1 is most relevant, and 0 is least relevant. Often, search results are sorted by score, but the client may specify a different sort order.
	 * See [Patient Match](patient-operation-match.html) for the EMPI search which relates to this Element.
	 */
	Score int `json:"score,omitempty"`
}

type BundleEntryRequest struct {
	/**
	 * Only perform the operation if the Etag value matches. For more information, see the API section ["Managing Resource Contention"](http.html#concurrency).
	 */
	IfMatch     string   `json:"ifMatch,omitempty"`
	Ext_ifMatch *Element `json:"_ifMatch,omitempty"`
	/**
	 * Only perform the operation if the last updated date matches. See the API documentation for ["Conditional Read"](http.html#cread).
	 */
	IfModifiedSince     string   `json:"ifModifiedSince,omitempty"`
	Ext_ifModifiedSince *Element `json:"_ifModifiedSince,omitempty"`
	/**
	 * Instruct the server not to perform the create if a specified resource already exists. For further information, see the API documentation for ["Conditional Create"](http.html#ccreate). This is just the query portion of the URL - what follows the "" (not including the "").
	 */
	IfNoneExist     string   `json:"ifNoneExist,omitempty"`
	Ext_ifNoneExist *Element `json:"_ifNoneExist,omitempty"`
	/**
	 * If the ETag values match, return a 304 Not Modified status. See the API documentation for ["Conditional Read"](http.html#cread).
	 */
	IfNoneMatch     string   `json:"ifNoneMatch,omitempty"`
	Ext_ifNoneMatch *Element `json:"_ifNoneMatch,omitempty"`
	/**
	 * In a transaction or batch, this is the HTTP action to be executed for this entry. In a history bundle, this indicates the HTTP action that occurred.
	 */
	Method     string   `json:"method,omitempty"`
	Ext_method *Element `json:"_method,omitempty"`
	/**
	 * E.g. for a Patient Create, the method would be "POST" and the URL would be "Patient". For a Patient Update, the method would be PUT and the URL would be "Patient/[id]".
	 */
	Url     string   `json:"url,omitempty"`
	Ext_url *Element `json:"_url,omitempty"`
}
type BundleEntryResponse struct {
	/**
	 * Etags match the Resource.meta.versionId. The ETag has to match the version id in the header if a resource is included.
	 */
	Etag     string   `json:"etag,omitempty"`
	Ext_etag *Element `json:"_etag,omitempty"`
	/**
	 * This has to match the same time in the meta header (meta.lastUpdated) if a resource is included.
	 */
	LastModified     string   `json:"lastModified,omitempty"`
	Ext_lastModified *Element `json:"_lastModified,omitempty"`
	/**
	 * The location header created by processing this operation, populated if the operation returns a location.
	 */
	Location     string   `json:"location,omitempty"`
	Ext_location *Element `json:"_location,omitempty"`
	/**
	 * For a POST/PUT operation, this is the equivalent outcome that would be returned for prefer = operationoutcome - except that the resource is always returned whether or not the outcome is returned.
	 * This outcome is not used for error responses in batch/transaction, only for hints and warnings. In a batch operation, the error will be in Bundle.entry.response, and for transaction, there will be a single OperationOutcome instead of a bundle in the case of an error.
	 */
	Outcome any `json:"outcome,omitempty"`
	/**
	 * The status code returned by processing this entry. The status SHALL start with a 3 digit HTTP code (e.g. 404) and may contain the standard HTTP description associated with the status code.
	 */
	Status     string   `json:"status,omitempty"`
	Ext_status *Element `json:"_status,omitempty"`
}

type Signature struct {
	/**
	 * Where the signature type is an XML DigSig, the signed content is a FHIR Resource(s), the signature is of the XML form of the Resource(s) using  XML-Signature (XMLDIG) "Detached Signature" form.
	 */
	Data     string   `json:"data,omitempty"`
	Ext_data *Element `json:"_data,omitempty"`
	/**
	 * The party that can't sign. For example a child.
	 */
	OnBehalfOf Reference `json:"onBehalfOf,omitempty"`
	/**
	 * A mime type that indicates the technical format of the signature. Important mime types are application/signature+xml for X ML DigSig, application/jose for JWS, and image/* for a graphical image of a signature, etc.
	 */
	SigFormat     string   `json:"sigFormat,omitempty"`
	Ext_sigFormat *Element `json:"_sigFormat,omitempty"`
	/**
	 * "xml", "json" and "ttl" are allowed, which describe the simple encodings described in the specification (and imply appropriate bundle support). Otherwise, mime types are legal here.
	 */
	TargetFormat     string   `json:"targetFormat,omitempty"`
	Ext_targetFormat *Element `json:"_targetFormat,omitempty"`
	/**
	 * Examples include attesting to authorship, correct transcription, and witness of specific event. Also known as a &quotCommitment Type Indication&quot.
	 */
	Type []Coding `json:"type,omitempty"`
	/**
	 * This should agree with the information in the signature.
	 */
	When     string   `json:"when,omitempty"`
	Ext_when *Element `json:"_when,omitempty"`
	/**
	 * This should agree with the information in the signature.
	 */
	Who *Reference `json:"who,omitempty"`
}

type Gender string

const (
	Male    Gender = "male"
	Female  Gender = "female"
	Other   Gender = "other"
	Unknown Gender = "unknown"
)

type BackboneElement struct {
	/**
	 * There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	 */
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}
type Element struct {
	/**
	 * There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	 */
	Extension []Extension `json:"extension,omitempty"`
	/**
	 * Unique id for the *Element within a resource (for internal references). This may be any string value that does not contain spaces.
	 */
	ID     string   `json:"id,omitempty"`
	Ext_id *Element `json:"_id,omitempty"`
}

type Resource struct {
	/** Resource Type Name (for serialization) */
	ResourceType string `json:"resourceType"`
	/**
	 * The only time that a resource does not have an id is when it is being submitted to the server using a create operation.
	 */
	ID     string   `json:"id,omitempty"`
	Ext_id *Element `json:"_id,omitempty"`
	/**
	 * Asserting this rule set restricts the content to be only understood by a limited set of trading partners. This inherently limits the usefulness of the data in the long term. However, the existing health eco-system is highly fractured, and not yet ready to define, collect, and exchange data in a generally computable sense. Wherever possible, implementers and/or specification writers should avoid using this *Element. Often, when used, the URL is a reference to an implementation guide that defines these special rules as part of it's narrative along with other profiles, value sets, etc.
	 */
	ImplicitRules     string   `json:"implicitRules,omitempty"`
	Ext_implicitRules *Element `json:"_implicitRules,omitempty"`
	/**
	 * Language is provided to support indexing and accessibility (typically, services such as text to speech use the language tag). The html language tag in the narrative applies  to the narrative. The language tag on the resource may be used to specify the language of other presentations generated from the data in the resource. Not all the content has to be in the base language. The Resource.language should not be assumed to apply to the narrative automatically. If a language is specified, it should it also be specified on the div *Element in the html (see rules in HTML5 for information about the relationship between xmllang and the html lang attribute).
	 */
	Language     string   `json:"language,omitempty"`
	Ext_language *Element `json:"_language,omitempty"`
	/**
	 * The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	 */
	Meta Meta `json:"meta,omitempty"`
}

type Meta struct {

	/**
	 * This value is always populated except when the resource is first being created. The server / resource manager sets this value what a client provides is irrelevant. This is equivalent to the HTTP Last-Modified and SHOULD have the same value on a [read](http.html#read) interaction.
	 */
	LastUpdated     string   `json:"lastUpdated,omitempty"`
	Ext_lastUpdated *Element `json:"Ext_LastUpdated,omitempty"`
	/**
	 * It is up to the server and/or other infrastructure of policy to determine whether/how these claims are verified and/or updated over time.  The list of profile URLs is a set.
	 */
	Profile     []string   `json:"profile,omitempty"`
	Ext_profile []*Element `json:"_profile,omitempty"`
	/**
	 * The security labels can be updated without changing the stated version of the resource. The list of security labels is a set. Uniqueness is based the system/code, and version and display are ignored.
	 */
	Security []Coding `json:"security,omitempty"`
	/**
	 * In the provenance resource, this corresponds to Provenance.entity.what[x]. The exact use of the source (and the implied Provenance.entity.role) is left to implementer discretion. Only one nominated source is allowed for additional provenance details, a full Provenance resource should be used.
	 * This *Element can be used to indicate where the current master source of a resource that has a canonical URL if the resource is no longer hosted at the canonical URL.
	 */
	Source     string   `json:"source,omitempty"`
	Ext_source *Element `json:"_source,omitempty"`
	/**
	 * The tags can be updated without changing the stated version of the resource. The list of tags is a set. Uniqueness is based the system/code, and version and display are ignored.
	 */
	Tag []Coding `json:"tag,omitempty"`
	/**
	 * The server assigns this value, and ignores what the client specifies, except in the case that the server is imposing version integrity on updates/deletes.
	 */
	VersionId     string   `json:"versionId,omitempty"`
	Ext_versionId *Element `json:"_versionId,omitempty"`
}

type Coding struct {
	/**
	 * A symbol in syntax defined by the system. The symbol may be a predefined code or an expression in a syntax defined by the coding system (e.g. post-coordination).
	 */
	Code     string   `json:"code,omitempty"`
	Ext_code *Element `json:"_code,omitempty"`
	/**
	 * A representation of the meaning of the code in the system, following the rules of the system.
	 */
	Display     string   `json:"display,omitempty"`
	Ext_display *Element `json:"_display,omitempty"`
	/**
	 * The URI may be an OID (urnoid...) or a UUID (urnuuid...).  OIDs and UUIDs SHALL be references to the HL7 OID registry. Otherwise, the URI should come from HL7's list of FHIR defined special URIs or it should reference to some definition that establishes the system clearly and unambiguously.
	 */
	System     string   `json:"system,omitempty"`
	Ext_system *Element `json:"_system,omitempty"`
	/**
	 * Amongst a set of alternatives, a directly chosen code is the most appropriate starting point for new translations. There is some ambiguity about what exactly 'directly chosen' implies, and trading partner agreement may be needed to clarify the use of this *Element and its consequences more completely.
	 */
	UserSelected     bool     `json:"userSelected,omitempty"`
	Ext_userSelected *Element `json:"_userSelected,omitempty"`
	/**
	 * Where the terminology does not clearly define what string should be used to identify code system versions, the recommendation is to use the date (expressed in FHIR date format) on which that version was officially published as the version date.
	 */
	Version     string   `json:"version,omitempty"`
	Ext_version *Element `json:"_version,omitempty"`
}

type Extension struct {

	/**
	 * The definition may point directly to a computable or human-readable definition of the extensibility codes, or it may be a logical URI as declared in some other specification. The definition SHALL be a URI for the Structure Definition defining the extension.
	 */
	Url     string   `json:"url,omitempty"`
	Ext_url *Element `json:"_url,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueBase64Binary     string   `json:"valueBase64Binary,omitempty"`
	Ext_valueBase64Binary *Element `json:"_valueBase64Binary,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueBoolean     bool     `json:"valueBoolean,omitempty"`
	Ext_valueBoolean *Element `json:"_valueBoolean,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueCanonical     string   `json:"valueCanonical,omitempty"`
	Ext_valueCanonical *Element `json:"_valueCanonical,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueCode     string   `json:"valueCode,omitempty"`
	Ext_valueCode *Element `json:"_valueCode,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueDate     string   `json:"valueDate,omitempty"`
	Ext_valueDate *Element `json:"_valueDate,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueDateTime     string   `json:"valueDateTime,omitempty"`
	Ext_valueDateTime *Element `json:"_valueDateTime,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueDecimal float64 `json:"valueDecimal,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueID     string   `json:"valueId,omitempty"`
	Ext_valueId *Element `json:"_valueId,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueInstant     string   `json:"valueInstant,omitempty"`
	Ext_valueInstant *Element `json:"_valueInstant,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueInteger int `json:"valueInteger,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueMarkdown     string   `json:"valueMarkdown,omitempty"`
	Ext_valueMarkdown *Element `json:"_valueMarkdown,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueOid     string   `json:"valueOid,omitempty"`
	Ext_valueOid *Element `json:"_valueOid,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValuePositiveInt int `json:"valuePositiveInt,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueString     string   `json:"valueString,omitempty"`
	Ext_valueString *Element `json:"_valueString,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueTime     string   `json:"valueTime,omitempty"`
	Ext_valueTime *Element `json:"_valueTime,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueUnsignedInt uint `json:"valueUnsignedInt,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueUri     string   `json:"valueUri,omitempty"`
	Ext_valueUri *Element `json:"_valueUri,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueUrl     string   `json:"valueUrl,omitempty"`
	Ext_valueUrl *Element `json:"_valueUrl,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */
	ValueUuid     string   `json:"valueUuid,omitempty"`
	Ext_valueUuid *Element `json:"_valueUuid,omitempty"`
	/**
	 * Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	 */

	ValueMeta *Meta `json:"valueMeta,omitempty"`
}

type Patient struct {

	/** Resource Type Name (for serialization) */
	ResourceType string `json:"resourceType""`
	/**
	 * If a record is inactive, and linked to an active record, then future patient/record updates should occur on the other patient.
	 */
	Active     bool     `json:"active,omitempty"`
	Ext_active *Element `json:"_active,omitempty"`
	/**
	 * Patient may have multiple addresses with different uses or applicable periods.
	 */
	Address []Address `json:"address,omitempty"`
	/**
	 * At least an estimated year should be provided as a guess if the real DOB is unknown  There is a standard extension "patient-birthTime" available that should be used where Time is required (such as in maternity/infant care systems).
	 */
	BirthDate     string   `json:"birthDate,omitempty"`
	Ext_birthDate *Element `json:"_birthDate,omitempty"`
	/**
	 * If no language is specified, this *implies* that the default local language is spoken.  If you need to convey proficiency for multiple modes, then you need multiple Patient.Communication associations.   For animals, language is not a relevant field, and should be absent from the instance. If the Patient does not speak the default local language, then the Interpreter Required Standard can be used to explicitly declare that an interpreter is required.
	 */
	Communication []PatientCommunication `json:"communication,omitempty"`
	/**
	 * Contact covers all kinds of contact parties family members, business contacts, guardians, caregivers. Not applicable to register pedigree and family ties beyond use of having contact.
	 */
	Contact []PatientContact `json:"contact,omitempty"`
	/**
	 * If there's no value in the instance, it means there is no statement on whether or not the individual is deceased. Most systems will interpret the absence of a value as a sign of the person being alive.
	 */
	DeceasedBoolean     bool     `json:"deceasedBoolean,omitempty"`
	Ext_deceasedBoolean *Element `json:"_deceasedBoolean,omitempty"`
	/**
	 * If there's no value in the instance, it means there is no statement on whether or not the individual is deceased. Most systems will interpret the absence of a value as a sign of the person being alive.
	 */
	DeceasedDateTime     string   `json:"deceasedDateTime,omitempty"`
	Ext_deceasedDateTime *Element `json:"_deceasedDateTime,omitempty"`
	/**
	 * The gender might not match the biological sex as determined by genetics or the individual's preferred identification. Note that for both humans and particularly animals, there are other legitimate possibilities than male and female, though the vast majority of systems and contexts only support male and female.  Systems providing decision support or enforcing business rules should ideally do this on the basis of Observations dealing with the specific sex or gender aspect of interest (anatomical, chromosomal, social, etc.)  However, because these observations are infrequently recorded, defaulting to the administrative gender is common practice.  Where such defaulting occurs, rule enforcement should allow for the variation between administrative and biological, chromosomal and other gender aspects.  For example, an alert about a hysterectomy on a male should be handled as a warning or overridable error, not a "hard" error.  See the Patient Gender and Sex section for additional information about communicating patient gender and sex.
	 */
	Gender     string   `json:"gender,omitempty"`
	Ext_gender *Element `json:"_gender,omitempty"`
	/**
	 * This may be the primary care provider (in a GP context), or it may be a patient nominated care manager in a community/disability setting, or even organization that will provide people to perform the care provider roles.  It is not to be used to record Care Teams, these should be in a CareTeam resource that may be linked to the CarePlan or EpisodeOfCare resources.
	 * Multiple GPs may be recorded against the patient for various reasons, such as a student that has his home GP listed along with the GP at university during the school semesters, or a "fly-in/fly-out" worker that has the onsite GP also included with his home GP to remain aware of medical issues.
	 * Jurisdictions may decide that they can profile this down to 1 if desired, or 1 per type.
	 */
	GeneralPractitioner []Reference `json:"generalPractitioner,omitempty"`
	/**
	 * An identifier for this patient.
	 */
	Identifier []Identifier `json:"identifier,omitempty"`
	/**
	 * There is no assumption that linked patient records have mutual links.
	 */
	Link []PatientLink `json:"link,omitempty"`
	/**
	 * There is only one managing organization for a specific patient record. Other organizations will have their own Patient record, and may use the Link property to join the records together (or a Person resource which can include confidence ratings for the association).
	 */
	ManagingOrganization *Reference `json:"managingOrganization,omitempty"`
	/**
	 * This field contains a patient's most recent marital (civil) status.
	 */
	MaritalStatus *CodeableConcept `json:"maritalStatus,omitempty"`
	/**
	 * Where the valueInteger is provided, the number is the birth number in the sequence. E.g. The middle birth in triplets would be valueInteger=2 and the third born would have valueInteger=3 If a boolean value was provided for this triplets example, then all 3 patient records would have valueBoolean=true (the ordering is not indicated).
	 */
	MultipleBirthBoolean     bool     `json:"multipleBirthBoolean,omitempty"`
	Ext_multipleBirthBoolean *Element `json:"_multipleBirthBoolean,omitempty"`
	/**
	 * Where the valueInteger is provided, the number is the birth number in the sequence. E.g. The middle birth in triplets would be valueInteger=2 and the third born would have valueInteger=3 If a boolean value was provided for this triplets example, then all 3 patient records would have valueBoolean=true (the ordering is not indicated).
	 */
	MultipleBirthInteger int `json:"multipleBirthInteger,omitempty"`
	/**
	 * A patient may have multiple names with different uses or applicable periods. For animals, the name is a "HumanName" in the sense that is assigned and used by humans and has the same patterns.
	 */
	Name []HumanName `json:"name,omitempty"`
	/**
	 * Guidelines
	 * * Use id photos, not clinical photos.
	 * * Limit dimensions to thumbnail.
	 * * Keep byte count low to ease resource updates.
	 */
	Photo []Attachment `json:"photo,omitempty"`
	/**
	 * A Patient may have multiple ways to be contacted with different uses or applicable periods.  May need to have options for contacting the person urgently and also to help with identification. The address might not go directly to the individual, but may reach another party that is able to proxy for the patient (i.e. home phone, or pet owner's phone).
	 */
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

type AddressType string

const (
	Postal   AddressType = "postal"
	Physical AddressType = "physical"
	Both     AddressType = "both"
)

type AddressUse string

const (
	Home    AddressUse = "home"
	Work    AddressUse = "work"
	Temp    AddressUse = "temp"
	Old     AddressUse = "old"
	Billing AddressUse = "billing"
)

type Address struct {
	/**
	 * The name of the city, town, suburb, village or other community or delivery center.
	 */
	City     string   `json:"city,omitempty"`
	Ext_city *Element `json:"_city,omitempty"`
	/**
	 * ISO 3166 3 letter codes can be used in place of a human readable country name.
	 */
	Country     string   `json:"country,omitempty"`
	Ext_country *Element `json:"_country,omitempty"`
	/**
	 * District is sometimes known as county, but in some regions 'county' is used in place of city (municipality), so county name should be conveyed in city instead.
	 */
	District     string   `json:"district,omitempty"`
	Ext_district *Element `json:"_district,omitempty"`
	/**
	 * This component contains the house number, apartment number, street name, street direction,  P.O. Box number, delivery hints, and similar address information.
	 */
	Line     []string  `json:"line,omitempty"`
	Ext_line []Element `json:"_line,omitempty"`
	/**
	 * Time period when address was/is in use.
	 */
	Period *Period `json:"period,omitempty"`
	/**
	 * A postal code designating a region defined by the postal service.
	 */
	PostalCode     string   `json:"postalCode,omitempty"`
	Ext_postalCode *Element `json:"_postalCode,omitempty"`
	/**
	 * Sub-unit of a country with limited sovereignty in a federally organized country. A code may be used if codes are in common use (e.g. US 2 letter state codes).
	 */
	State     string   `json:"state,omitempty"`
	Ext_state *Element `json:"_state,omitempty"`
	/**
	 * Can provide both a text representation and parts. Applications updating an address SHALL ensure that  when both text and parts are present,  no content is included in the text that isn't found in a part.
	 */
	Text     string   `json:"text,omitempty"`
	Ext_text *Element `json:"_text,omitempty"`
	/**
	 * The definition of Address states that "address is intended to describe postal addresses, not physical locations". However, many applications track whether an address has a dual purpose of being a location that can be visited as well as being a valid delivery destination, and Postal addresses are often used as proxies for physical locations (also see the [Location](location.html#) resource).
	 */
	Type     string   `json:"type,omitempty"`
	Ext_type *Element `json:"_type,omitempty"`
	/**
	 * Applications can assume that an address is current unless it explicitly says that it is temporary or old.
	 */
	Use     string   `json:"use,omitempty"`
	Ext_use *Element `json:"_use,omitempty"`
}

type Period struct {
	/**
	 * The high value includes any matching date/time. i.e. 2012-02-03T100000 is in a period that has an end value of 2012-02-03.
	 */
	End     string   `json:"end,omitempty"`
	Ext_end *Element `json:"_end,omitempty"`
	/**
	 * If the low *Element is missing, the meaning is that the low boundary is not known.
	 */
	Start     string   `json:"start,omitempty"`
	Ext_start *Element `json:"_start,omitempty"`
}

type PatientCommunication struct {
	/**
	 * The structure aa-BB with this exact casing is one the most widely used notations for locale. However not all systems actually code this but instead have it as free text. Hence CodeableConcept instead of code as the data type.
	 */
	Language *CodeableConcept `json:"language,omitempty"`
	/**
	 * This language is specifically identified for communicating healthcare information.
	 */
	Preferred     bool     `json:"preferred,omitempty"`
	Ext_preferred *Element `json:"_preferred,omitempty"`
}

type CodeableConcept struct {
	/**
	 * Codes may be defined very casually in enumerations, or code lists, up to very formal definitions such as SNOMED CT - see the HL7 v3 Core Principles for more information.  Ordering of codings is undefined and SHALL NOT be used to infer meaning. Generally, at most only one of the coding values will be labeled as UserSelected = true.
	 */
	Coding []Coding `json:"coding,omitempty"`
	/**
	 * Very often the text is the same as a displayName of one of the codings.
	 */
	Text     string   `json:"text,omitempty"`
	Ext_text *Element `json:"_text,omitempty"`
}

type PatientContact struct {
	/**
	 * Address for the contact person.
	 */
	Address *Address `json:"address,omitempty"`
	/**
	 * Administrative Gender - the gender that the contact person is considered to have for administration and record keeping purposes.
	 */
	Gender     string   `json:"gender,omitempty"`
	Ext_gender *Element `json:"_gender,omitempty"`
	/**
	 * A name associated with the contact person.
	 */
	Name *HumanName `json:"name,omitempty"`
	/**
	 * Organization on behalf of which the contact is acting or for which the contact is working.
	 */
	Organization *Reference `json:"organization,omitempty"`
	/**
	 * The period during which this contact person or organization is valid to be contacted relating to this patient.
	 */
	Period *Period `json:"period,omitempty"`
	/**
	 * The nature of the relationship between the patient and the contact person.
	 */
	Relationship []CodeableConcept `json:"relationship,omitempty"`
	/**
	 * Contact may have multiple ways to be contacted with different uses or applicable periods.  May need to have options for contacting the person urgently, and also to help with identification.
	 */
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

type HumanNameUse string

const (
	Usual     HumanNameUse = "usual"
	Official  HumanNameUse = "official"
	HNTemp    HumanNameUse = "temp"
	Nickname  HumanNameUse = "nickname"
	Anonymous HumanNameUse = "anonymous"
	HNOld     HumanNameUse = "old"
	Maiden    HumanNameUse = "maiden"
)

type HumanName struct {
	/**
	 * Family Name may be decomposed into specific parts using extensions (de, nl, es related cultures).
	 */
	Family     string   `json:"family,omitempty"`
	Ext_family *Element `json:"_family,omitempty"`
	/**
	 * If only initials are recorded, they may be used in place of the full name parts. Initials may be separated into multiple given names but often aren't due to paractical limitations.  This *Element is not called "first name" since given names do not always come first.
	 */
	Given     []string  `json:"given,omitempty"`
	Ext_given []Element `json:"_given,omitempty"`
	/**
	 * Indicates the period of time when this name was valid for the named person.
	 */
	Period *Period `json:"period,omitempty"`
	/**
	 * Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.
	 */
	Prefix     []string  `json:"prefix,omitempty"`
	Ext_prefix []Element `json:"_prefix,omitempty"`
	/**
	 * Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.
	 */
	Suffix     []string  `json:"suffix,omitempty"`
	Ext_suffix []Element `json:"_suffix,omitempty"`
	/**
	 * Can provide both a text representation and parts. Applications updating a name SHALL ensure that when both text and parts are present,  no content is included in the text that isn't found in a part.
	 */
	Text     string   `json:"text,omitempty"`
	Ext_text *Element `json:"_text,omitempty"`
	/**
	 * Applications can assume that a name is current unless it explicitly says that it is temporary or old.
	 */
	Use     string   `json:"use,omitempty"`
	Ext_use *Element `json:"_use,omitempty"`
}

type IdentifierUse string

const (
	IUsual    IdentifierUse = "usual"
	IOfficial IdentifierUse = "official"
	ITemp     IdentifierUse = "temp"
	Secondary IdentifierUse = "secondary"
	IOld      IdentifierUse = "old"
)

type Identifier struct {
	/**
	 * The Identifier.assigner may omit the .reference *Element and only contain a .display *Element reflecting the name or other textual information about the assigning organization.
	 */
	Assigner *Reference `json:"assigner,omitempty"`
	/**
	 * Time period during which identifier is/was valid for use.
	 */
	Period *Period `json:"period,omitempty"`
	/**
	 * Identifier.system is always case sensitive.
	 */
	System     string   `json:"system,omitempty"`
	Ext_system *Element `json:"_system,omitempty"`
	/**
	 * This *Element deals only with general categories of identifiers.  It SHOULD not be used for codes that correspond 1..1 with the Identifier.system. Some identifiers may fall into multiple categories due to common usage.   Where the system is known, a type is unnecessary because the type is always part of the system definition. However systems often need to handle identifiers where the system is not known. There is not a 11 relationship between type and system, since many different systems have the same type.
	 */
	Type *CodeableConcept `json:"type,omitempty"`
	/**
	 * Applications can assume that an identifier is permanent unless it explicitly says that it is temporary.
	 */
	Use     string   `json:"use,omitempty"`
	Ext_use *Element `json:"_use,omitempty"`
	/**
	 * If the value is a full URI, then the system SHALL be urnietfrfc3986.  The value's primary purpose is computational mapping.  As a result, it may be normalized for comparison purposes (e.g. removing non-significant whitespace, dashes, etc.)  A value formatted for human display can be conveyed using the [Rendered Value extension](extension-rendered-value.html). Identifier.value is to be treated as case sensitive unless knowledge of the Identifier.system allows the processer to be confident that non-case-sensitive processing is safe.
	 */
	Value     string   `json:"value,omitempty"`
	Ext_value *Element `json:"_value,omitempty"`
}

type Reference struct {
	/**
	 * This is generally not the same as the Resource.text of the referenced resource.  The purpose is to identify what's being referenced, not to fully describe it.
	 */
	Display     string   `json:"display,omitempty"`
	Ext_display *Element `json:"_display,omitempty"`
	/**
	 * When an identifier is provided in place of a reference, any system processing the reference will only be able to resolve the identifier to a reference if it understands the business context in which the identifier is used. Sometimes this is global (e.g. a national identifier) but often it is not. For this reason, none of the useful mechanisms described for working with references (e.g. chaining, includes) are possible, nor should servers be expected to be able resolve the reference. Servers may accept an identifier based reference untouched, resolve it, and/or reject it - see CapabilityStatement.rest.resource.referencePolicy.
	 * When both an identifier and a literal reference are provided, the literal reference is preferred. Applications processing the resource are allowed - but not required - to check that the identifier matches the literal reference
	 * Applications converting a logical reference to a literal reference may choose to leave the logical reference present, or remove it.
	 * Reference is intended to point to a structure that can potentially be expressed as a FHIR resource, though there is no need for it to exist as an actual FHIR resource instance - except in as much as an application wishes to actual find the target of the reference. The content referred to be the identifier must meet the logical constraints implied by any limitations on what resource types are permitted for the reference.  For example, it would not be legitimate to send the identifier for a drug prescription if the type were Reference(Observation|DiagnosticReport).  One of the use-cases for Reference.identifier is the situation where no FHIR representation exists (where the type is Reference (Any).
	 */
	Identifier *Identifier `json:"identifier,omitempty"`
	/**
	 * Using absolute URLs provides a stable scalable approach suitable for a cloud/web context, while using relative/logical references provides a flexible approach suitable for use when trading across closed eco-system boundaries.   Absolute URLs do not need to point to a FHIR RESTful server, though this is the preferred approach. If the URL conforms to the structure "/[type]/[id]" then it should be assumed that the reference is to a FHIR RESTful server.
	 */
	Reference     string   `json:"reference,omitempty"`
	Ext_reference *Element `json:"_reference,omitempty"`
	/**
	 * This *Element is used to indicate the type of  the target of the reference. This may be used which ever of the other *Elements are populated (or not). In some cases, the type of the target may be determined by inspection of the reference (e.g. a RESTful URL) or by resolving the target of the reference if both the type and a reference is provided, the reference SHALL resolve to a resource of the same type as that specified.
	 */
	Type     string   `json:"type,omitempty"`
	Ext_type *Element `json:"_type,omitempty"`
}

type ContactPointSystem string

const (
	Phone   ContactPointSystem = "phone"
	Fax     ContactPointSystem = "fax"
	Email   ContactPointSystem = "email"
	Pager   ContactPointSystem = "pager"
	URL     ContactPointSystem = "url"
	SMS     ContactPointSystem = "sms"
	CPOther ContactPointSystem = "other"
)

type ContactPointUse string

const (
	CPHome   ContactPointUse = "home"
	CPWork   ContactPointUse = "work"
	CPTemp   ContactPointUse = "temp"
	CPOld    ContactPointUse = "old"
	CPMobile ContactPointUse = "mobile"
)

type ContactPoint struct {
	/**
	 * Time period when the contact point was/is in use.
	 */
	Period *Period `json:"period,omitempty"`
	/**
	 * Note that rank does not necessarily follow the order in which the contacts are represented in the instance.
	 */
	rank int `json:"rank,omitempty"`
	/**
	 * Telecommunications form for contact point - what communications system is required to make use of the contact.
	 */
	System     string   `json:"system,omitempty"`
	Ext_system *Element `json:"_system,omitempty"`
	/**
	 * Applications can assume that a contact is current unless it explicitly says that it is temporary or old.
	 */
	Use     string   `json:"use,omitempty"`
	Ext_use *Element `json:"_use,omitempty"`
	/**
	 * Additional text data such as phone extension numbers, or notes about use of the contact are sometimes included in the value.
	 */
	Value     string   `json:"value,omitempty"`
	Ext_value *Element `json:"_value,omitempty"`
}

type Attachment struct {

	/**
	 * Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. Includes mime type parameters such as charset where appropriate.
	 */
	ContentType     string   `json:"contentType,omitempty"`
	Ext_contentType *Element `json:"_contentType,omitempty"`
	/**
	 * The date that the attachment was first created.
	 */
	Creation     string   `json:"creation,omitempty"`
	Ext_creation *Element `json:"_creation,omitempty"`
	/**
	 * The base64-encoded data SHALL be expressed in the same character set as the base resource XML or JSON.
	 */
	Data     string   `json:"data,omitempty"`
	Ext_data *Element `json:"_data,omitempty"`
	/**
	 * The hash is calculated on the data prior to base64 encoding, if the data is based64 encoded. The hash is not intended to support digital signatures. Where protection against malicious threats a digital signature should be considered, see [Provenance.signature](provenance-definitions.html#Provenance.signature) for mechanism to protect a resource with a digital signature.
	 */
	Hash     string   `json:"hash,omitempty"`
	Ext_hash *Element `json:"_hash,omitempty"`
	/**
	 * The human language of the content. The value can be any valid value according to BCP 47.
	 */
	Language     string   `json:"language,omitempty"`
	Ext_language *Element `json:"_language,omitempty"`
	/**
	 * The number of bytes is redundant if the data is provided as a base64binary, but is useful if the data is provided as a url reference.
	 */
	Size int `json:"size,omitempty"`
	/**
	 * A label or set of text to display in place of the data.
	 */
	Title     string   `json:"title,omitempty"`
	Ext_title *Element `json:"_title,omitempty"`
	/**
	 * If both data and url are provided, the url SHALL point to the same content as the data contains. Urls may be relative references or may reference transient locations such as a wrapping envelope using cid though this has ramifications for using signatures. Relative URLs are interpreted relative to the service url, like a resource reference, rather than relative to the resource itself. If a URL is provided, it SHALL resolve to actual data.
	 */
	Url     string   `json:"url,omitempty"`
	Ext_url *Element `json:"_url,omitempty"`
}
type PatientLinkType string

const (
	ReplacedBy PatientLinkType = "replaced-by"
	Replaces   PatientLinkType = "Replaces"
	Refer      PatientLinkType = "refer"
	SeeAlso    PatientLinkType = "seealso"
)

type PatientLink struct {
	BackboneElement `json:"backboneElement,omitempty"`
	/**
	 * Referencing a RelatedPerson here removes the need to use a Person record to associate a Patient and RelatedPerson as the same individual.
	 */
	Other *Reference `json:"other,omitempty"`
	/**
	 * The type of link between this patient resource and another patient resource.
	 */
	Type     PatientLinkType `json:"type,omitempty"`
	Ext_type *Element        `json:"_type,omitempty"`
}
type MedicationStatus string

const (
	Active         MedicationStatus = "active"
	Inactive       MedicationStatus = "inactive"
	EnteredInError MedicationStatus = "entered-in-error"
)

type Medication struct {

	/** Resource Type Name (for serialization) */
	ResourceType string `json:"resourceType"`
	/**
	 * Specific amount of the drug in the packaged product.  For example, when specifying a product that has the same strength (For example, Insulin glargine 100 unit per mL solution for injection), this attribute provides additional clarification of the package amount (For example, 3 mL, 10mL, etc.).
	 */
	Amount *Ratio `json:"amount,omitempty"`
	/**
	 * Information that only applies to packages (not products).
	 */
	Batch *MedicationBatch `json:"batch,omitempty"`
	/**
	 * Depending on the context of use, the code that was actually selected by the user (prescriber, dispenser, etc.) will have the coding.userSelected set to true.  As described in the coding datatype: "A coding may be marked as a "userSelected" if a user selected the particular coded value in a user interface (e.g. the user selects an item in a pick-list). If a user selected coding exists, it is the preferred choice for performing translations etc. Other codes can only be literal translations to alternative code systems, or codes at a lower level of granularity (e.g. a generic code for a vendor-specific primary one).
	 */
	Code *CodeableConcept `json:"code,omitempty"`
	/**
	 * When Medication is referenced from MedicationRequest, this is the ordered form.  When Medication is referenced within MedicationDispense, this is the dispensed form.  When Medication is referenced within MedicationAdministration, this is administered form.
	 */
	Form *CodeableConcept `json:"form,omitempty"`
	/**
	 * The serial number could be included as an identifier.
	 */
	Identifier []Identifier `json:"identifier,omitempty"`
	/**
	 * The ingredients need not be a complete list.  If an ingredient is not specified, this does not indicate whether an ingredient is present or absent.  If an ingredient is specified it does not mean that all ingredients are specified.  It is possible to specify both inactive and active ingredients.
	 */
	Ingredient []MedicationIngredient `json:"ingredient,omitempty"`
	/**
	 * Describes the details of the manufacturer of the medication product.  This is not intended to represent the distributor of a medication product.
	 */
	Manufacturer *Reference `json:"manufacturer,omitempty"`
	/**
	 * This status is intended to identify if the medication in a local system is in active use within a drug database or inventory.  For example, a pharmacy system may create a new drug file record for a compounded product "ABC Hospital Special Cream" with an active status.  At some point in the future, it may be determined that the drug record was created with an error and the status is changed to "entered in error".   This status is not intended to specify if a medication is part of a particular formulary.  It is possible that the drug record may be referenced by multiple formularies or catalogues and each of those entries would have a separate status.
	 */
	Status     string   `json:"status,omitempty"`
	Ext_status *Element `json:"_status,omitempty"`
}

type Ratio struct {
	/**
	 * The value of the denominator.
	 */
	Denominator *Quantity `json:"denominator,omitempty"`
	/**
	 * The value of the numerator.
	 */
	Numerator *Quantity `json:"numerator,omitempty"`
}

type QuantityComparator string

const (
	Less        QuantityComparator = "<"
	LessEqual   QuantityComparator = "<="
	Grater      QuantityComparator = ">"
	GraterEqual QuantityComparator = ">="
)

type Quantity struct {
	/**
	 * The preferred system is UCUM, but SNOMED CT can also be used (for customary units) or ISO 4217 for currency.  The context of use may additionally require a code from a particular system.
	 */
	Code     string   `json:"code,omitempty"`
	Ext_code *Element `json:"_code,omitempty"`
	/**
	 * How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues e.g. if the comparator is "<" , then the real value is < stated value.
	 */
	Comparator     QuantityComparator `json:"comparator,omitempty"`
	Ext_comparator *Element           `json:"_comparator,omitempty"`
	/**
	 * The identification of the system that provides the coded form of the unit.
	 */
	System     string   `json:"system,omitempty"`
	Ext_system *Element `json:"_system,omitempty"`
	/**
	 * A human-readable form of the unit.
	 */
	Unit     string   `json:"unit,omitempty"`
	Ext_unit *Element `json:"_unit,omitempty"`
	/**
	 * The implicit precision in the value should always be honored. Monetary values have their own rules for handling precision (refer to standard accounting text books).
	 */
	Value int `json:"value,omitempty"`
}

type MedicationBatch struct {
	/**
	 * When this specific batch of product will expire.
	 */
	ExpirationDate     string   `json:"expirationDate,omitempty"`
	Ext_expirationDate *Element `json:"_expirationDate,omitempty"`
	/**
	 * The assigned lot number of a batch of the specified product.
	 */
	LotNumber     string   `json:"lotNumber,omitempty"`
	Ext_lotNumber *Element `json:"_lotNumber,omitempty"`
}

type MedicationIngredient struct {
	/**
	 * Indication of whether this ingredient affects the therapeutic action of the drug.
	 */
	IsActive     bool     `json:"isActive,omitempty"`
	Ext_isActive *Element `json:"_isActive,omitempty"`
	/**
	 * The actual ingredient - either a substance (simple ingredient) or another medication of a medication.
	 */
	Item ItemCodeableConcept `json:"item,omitempty"`
	/**
	 * The actual ingredient - either a substance (simple ingredient) or another medication of a medication.
	 */
	ItemReference *Reference `json:"itemReference,omitempty"`
	/**
	 * Specifies how many (or how much) of the items there are in this Medication.  For example, 250 mg per tablet.  This is expressed as a ratio where the numerator is 250mg and the denominator is 1 tablet.
	 */
	Strength *Ratio `json:"strength,omitempty"`
}
type ItemCodeableConcept struct {
	Concept *CodeableConcept `json:"concept,omitempty"`
}
