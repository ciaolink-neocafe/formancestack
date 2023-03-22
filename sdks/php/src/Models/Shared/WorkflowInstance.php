<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


class WorkflowInstance
{
	#[\JMS\Serializer\Annotation\SerializedName('createdAt')]
    #[\JMS\Serializer\Annotation\Type("DateTime<'Y-m-d\TH:i:s.up'>")]
    public \DateTime $createdAt;
    
	#[\JMS\Serializer\Annotation\SerializedName('error')]
    #[\JMS\Serializer\Annotation\Type('string')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?string $error = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('id')]
    #[\JMS\Serializer\Annotation\Type('string')]
    public string $id;
    
    /**
     * $status
     * 
     * @var ?array<\formance\stack\Models\Shared\StageStatus> $status
     */
	#[\JMS\Serializer\Annotation\SerializedName('status')]
    #[\JMS\Serializer\Annotation\Type('array<formance\stack\Models\Shared\StageStatus>')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?array $status = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('terminated')]
    #[\JMS\Serializer\Annotation\Type('bool')]
    public bool $terminated;
    
	#[\JMS\Serializer\Annotation\SerializedName('terminatedAt')]
    #[\JMS\Serializer\Annotation\Type("DateTime<'Y-m-d\TH:i:s.up'>")]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?\DateTime $terminatedAt = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('updatedAt')]
    #[\JMS\Serializer\Annotation\Type("DateTime<'Y-m-d\TH:i:s.up'>")]
    public \DateTime $updatedAt;
    
	#[\JMS\Serializer\Annotation\SerializedName('workflowID')]
    #[\JMS\Serializer\Annotation\Type('string')]
    public string $workflowID;
    
	public function __construct()
	{
		$this->createdAt = new \DateTime();
		$this->error = null;
		$this->id = "";
		$this->status = null;
		$this->terminated = false;
		$this->terminatedAt = null;
		$this->updatedAt = new \DateTime();
		$this->workflowID = "";
	}
}
