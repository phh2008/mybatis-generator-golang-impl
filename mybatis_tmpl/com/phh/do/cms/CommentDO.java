package com.phh.do.cms;

import java.util.Date;

import com.phh.do.BaseDO;

/**
 * 评论
 * 表名：cms_comment
 *
 * @author 彭会会
 * @date 2019-08-12
 */
public class CommentDO extends BaseDO {
	private static final long serialVersionUID = 1565586068591792400L;


    /**
     * 评论用户ID
     */
    private Long userId;

    /**
     * 评论用户昵称
     */
    private String nickName;

    /**
     * 被回复用户ID
     */
    private Long replyUserId;

    /**
     * 被回复用户呢称
     */
    private String replyNickName;

    /**
     * 文章ID
     */
    private Long articleId;

    /**
     * 评论内容
     */
    private String content;

    /**
     * 创建时间
     */
    private Date createdAt;

    /**
     * 更新时间
     */
    private Date updatedAt;

    /**
     * 删除状态：1-正常，2-逻辑删除，3-物理删除
     */
    private Integer state;



    /**
     * 评论用户ID
     */
    public void setUserId(Long userId) {
        this.userId = userId;
    }
    /**
     * 评论用户ID
     */
    public Long getUserId() {
        return userId;
    }

    /**
     * 评论用户昵称
     */
    public void setNickName(String nickName) {
        this.nickName = nickName;
    }
    /**
     * 评论用户昵称
     */
    public String getNickName() {
        return nickName;
    }

    /**
     * 被回复用户ID
     */
    public void setReplyUserId(Long replyUserId) {
        this.replyUserId = replyUserId;
    }
    /**
     * 被回复用户ID
     */
    public Long getReplyUserId() {
        return replyUserId;
    }

    /**
     * 被回复用户呢称
     */
    public void setReplyNickName(String replyNickName) {
        this.replyNickName = replyNickName;
    }
    /**
     * 被回复用户呢称
     */
    public String getReplyNickName() {
        return replyNickName;
    }

    /**
     * 文章ID
     */
    public void setArticleId(Long articleId) {
        this.articleId = articleId;
    }
    /**
     * 文章ID
     */
    public Long getArticleId() {
        return articleId;
    }

    /**
     * 评论内容
     */
    public void setContent(String content) {
        this.content = content;
    }
    /**
     * 评论内容
     */
    public String getContent() {
        return content;
    }

    /**
     * 创建时间
     */
    public void setCreatedAt(Date createdAt) {
        this.createdAt = createdAt;
    }
    /**
     * 创建时间
     */
    public Date getCreatedAt() {
        return createdAt;
    }

    /**
     * 更新时间
     */
    public void setUpdatedAt(Date updatedAt) {
        this.updatedAt = updatedAt;
    }
    /**
     * 更新时间
     */
    public Date getUpdatedAt() {
        return updatedAt;
    }

    /**
     * 删除状态：1-正常，2-逻辑删除，3-物理删除
     */
    public void setState(Integer state) {
        this.state = state;
    }
    /**
     * 删除状态：1-正常，2-逻辑删除，3-物理删除
     */
    public Integer getState() {
        return state;
    }


}
