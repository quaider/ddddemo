package cargo

import (
	"errors"
	"go-ddd/domain/handling"
	"go-ddd/domain/location"
)

// Cargo
/*
表示货运聚合的聚合根，是领域模型的核心。

在运输过程中，可以根据客户的要求重新安排货物的路线；
如果为货物指定了新路线，并请求了新路线。作为值对象的 Itinerary 将被丢弃，并附加新的 Itinerary。
当然，还可能发生货物意外错线的情况，这应通知适当的人员，并触发re-route

处理货物时，运送状态(delivery status)会发生变化，有关货物运送的所有内容都包含在 Delivery 值对象中，每当货物被由处理事件注册触发的异步事件处理时，该对象就会被替换。
当货物的路线被重新安排时，在cargo聚合中需同步更新货物运送数据
*/
type Cargo struct {
	trackingId         *TrackingId         // 货物跟踪id, 唯一标识货物
	origin             *location.Location  // 货物起始地，永不改变，初始从route specification中获取
	routeSpecification *RouteSpecification // 路径规格，表明 货运的起始地、目的地、截止到达时间
	itinerary          *Itinerary          // 货运航线，包含多个航程
	delivery           *Delivery           // 货运历史
}

func NewCargo(trackingId *TrackingId, route *RouteSpecification) (*Cargo, error) {
	if trackingId == nil {
		return nil, errors.New("tracking id is required")
	}

	if route == nil {
		return nil, errors.New("route specification is required")
	}

	c := &Cargo{
		trackingId:         trackingId,
		routeSpecification: route,
		origin:             route.Origin(),
	}

	delivery, err := DerivedFrom(route, nil, handling.EmptyHistory)
	if err != nil {
		return nil, err
	}

	c.delivery = delivery

	return c, nil
}

func (c *Cargo) Origin() *location.Location {
	return c.origin
}

func (c *Cargo) RouteSpecification() *RouteSpecification {
	return c.routeSpecification
}

func (c *Cargo) Itinerary() *Itinerary {
	return c.itinerary
}

func (c *Cargo) Delivery() *Delivery {
	return c.delivery
}

func (c *Cargo) TrackingId() *TrackingId {
	return c.trackingId
}

func (c *Cargo) SpecifyNewRoute(route *RouteSpecification) error {
	if route == nil {
		return errors.New("tracking id is required")
	}

	c.routeSpecification = route

	// 处理航线和货运数据
	d, err := c.delivery.UpdateOnRouting(route, c.itinerary)
	if err != nil {
		return err
	}

	c.delivery = d

	return nil
}

// AssignToRoute 为 Cargo 指定新航线
func (c *Cargo) AssignToRoute(newItinerary *Itinerary) error {
	if newItinerary == nil {
		return errors.New("itinerary is required for assignment")
	}

	c.itinerary = newItinerary

	d, err := c.delivery.UpdateOnRouting(c.routeSpecification, c.itinerary)
	if err != nil {
		return err
	}

	c.delivery = d

	return nil
}

func (c *Cargo) DeriveDeliveryProgress(h *handling.History) error {
	// Delivery is a value object, so we can simply discard the old one and replace it with a new
	d, err := DerivedFrom(c.routeSpecification, c.itinerary, h)
	if err != nil {
		return err
	}

	c.delivery = d

	return nil
}
